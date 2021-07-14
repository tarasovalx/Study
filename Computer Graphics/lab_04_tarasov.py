import numpy as np
import pyglet
from numba import jit
from pyglet.gl import *
from pyglet.window.key import *

polygon_vertices = []

window = pyglet.window.Window(width=400, height=400, resizable=True)
window.set_minimum_size(300, 300)
MAX_W, MAX_H = 1000, 1000
window.set_maximum_size(MAX_W, MAX_H)

fill_mode = False


def calc_intersection(y, edge):
    return int((edge.p1[0] - edge.p2[0]) * (y - edge.p2[1]) / (edge.p1[1] - edge.p2[1]) + edge.p2[0])


class Edge:
    def __init__(self, p1, p2):
        self.p1 = p1 if p1[1] > p2[1] else p2
        self.p2 = p1 if p1[1] <= p2[1] else p2
        self.y_max = self.p1[1]
        self.y_min = self.p2[1]
        self.x_min = min(p1[0], p2[0])
        self.x_max = max(p1[0], p2[0])
        self.x = self.p1[0]

    def __repr__(self):
        return f'(y_max: {self.y_max}, x: {self.x})'

    def intersec(self, y):
        self.x = calc_intersection(y, self)


c = 2
frame_buffer = np.zeros(shape=(MAX_W * c + 1, MAX_H * c + 1), dtype=np.float32)


def make_edges_list(l):
    verts = l.copy()
    for i in range(len(verts)):
        verts[i] = verts[i] * c

    edges = []
    for i in range(len(verts) - 1):
        if (verts[i][1] == verts[i + 1][1]):
            continue
        edges.append(Edge(verts[i], verts[i + 1]))
    edges.append(Edge(verts[0], verts[len(verts) - 1]))
    edges.sort(key=lambda e: e.y_max, reverse=True)
    return edges


def rasterize_polygon(dest_buffer):
    edges = make_edges_list(polygon_vertices)
    active_edges = []
    edge_p = 0
    for scan_line in range(window.height * 2, -1, -1):
        # Перебираем ещё недобавленные ребра
        for j in range(edge_p, len(edges)):
            if edges[j].y_max == scan_line:
                active_edges.append(edges[j])
                edge_p += 1
            else:
                break
        j = 0

        # Перебираем список активных ребер
        while j < len(active_edges):
            if active_edges[j].y_min > scan_line:
                active_edges.pop(j)
                continue
            active_edges[j].intersec(scan_line)
            j += 1
        active_edges.sort(key=lambda e: e.x)

        i = 1

        # Проход по точкам пересечения
        while i < len(active_edges):
            l_x, r_x = active_edges[i - 1].x, active_edges[i].x

            if i > 0 and r_x == l_x:
                if (active_edges[i].y_max - scan_line > 0) != (active_edges[i - 1].y_max - scan_line > 0):
                    i += 1
                    continue

            dest_buffer[l_x: r_x + 1, scan_line].fill(1.0)

            if (i < len(active_edges) - 1) and active_edges[i].x == active_edges[i + 1].x:
                if (active_edges[i].y_max - scan_line > 0) != (active_edges[i + 1].y_max - scan_line > 0):
                    i += 1

            i += 2


@jit
def apply_filter(buffer, kernel=np.ones((3, 3))):
    b_shape = buffer.shape
    for i in range(1, b_shape[0] - 1, 2):
        for j in range(1, b_shape[1] - 1, 2):
            buffer[i // c, j // c] = buffer[i - 1:i + 2, j - 1:j + 2].sum() / 9


@window.event
def on_key_press(symbol, modifiers):
    global polygon_vertices, frame_buffer, fill_mode

    if symbol == SPACE:
        frame_buffer.fill(0)
        fill_mode = False
        polygon_vertices.clear()

    if symbol == G and len(polygon_vertices) > 2:
        frame_buffer.fill(0)
        rasterize_polygon(frame_buffer)
        apply_filter(frame_buffer)
        fill_mode = True

    return pyglet.event.EVENT_HANDLED


@window.event
def on_mouse_release(x, y, button, modifiers):
    polygon_vertices.append(np.array([x, y], dtype=int))
    return pyglet.event.EVENT_HANDLED


@window.event
def on_resize(width, height):
    global frame_buffer
    frame_buffer.fill(0.0)
    if len(polygon_vertices) > 2:
        rasterize_polygon(frame_buffer)
        apply_filter(frame_buffer)
    glViewport(0, 0, width, height)

    return pyglet.event.EVENT_HANDLED


@window.event
def on_draw():
    window.clear()
    glMatrixMode(gl.GL_PROJECTION)
    glLoadIdentity()
    glOrtho(0, window.width, 0, window.height, 0, 1)

    glMatrixMode(gl.GL_MODELVIEW)
    glDrawPixels(window.width, window.height, GL_BLUE, GL_FLOAT,
                 frame_buffer[:window.width, :window.height].ravel('F').ctypes)

    if not fill_mode:
        glBegin(GL_LINE_STRIP)
        for vert in polygon_vertices:
            glVertex2f(*vert)

        if len(polygon_vertices) > 0:
            glVertex2f(*polygon_vertices[0])
        glEnd()


pyglet.app.run()
