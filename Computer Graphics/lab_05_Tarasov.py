import numpy as np
import pyglet
from pyglet import gl
from pyglet.gl import *
from pyglet.window.key import *

window = pyglet.window.Window(width=400, height=400, resizable=True)
window.set_minimum_size(300, 300)
window.set_maximum_size(1000, 1000)

cutter_vertices , lines_vertices = [], []

cutter_edges, lines = [], []

clipped_lines = []


class Edge:
    def __init__(self, p1, p2):
        self.p1 = p1
        self.p2 = p2
        self.N = np.array([p1[1] - p2[1], p2[0] - p1[0]])

    def __repr__(self):
        return f'({self.p1}, {self.p2}, {self.N})'


class Line:
    def __init__(self, p1, p2):
        self.p2 = p2 if p2[1] > p1[1] else p1
        self.p1 = p2 if p2[1] <= p1[1] else p1
        self.D = self.p2 - self.p1
        self.t = 0
        self.t_min = 0
        self.t_max = 1


@window.event
def on_mouse_release(x, y, button, modifiers):
    if button == 1:
        lines_vertices.append(np.array([x, y], dtype=float))
    elif button == 4:
        cutter_vertices.append(np.array([x, y], dtype=float))

    return pyglet.event.EVENT_HANDLED


@window.event
def on_key_press(symbol, modifiers):
    if symbol == SPACE:
        lines_vertices.clear()
        cutter_vertices.clear()
        lines.clear()
        clipped_lines.clear()
        cutter_edges.clear()
        pass

    if symbol == G:
        cutter_edges.clear()
        for i in range(len(cutter_vertices) - 1):
            cutter_edges.append(Edge(cutter_vertices[i], cutter_vertices[i + 1]))
        cutter_edges.append(Edge(cutter_vertices[len(cutter_vertices) - 1], cutter_vertices[0]))

        lines.clear()
        clipped_lines.clear()
        for i in range(0, len(lines_vertices) - 1, 2):
            lines.append(Line(lines_vertices[i], lines_vertices[i + 1]))

        for line in lines:
            clipped = clip_line(line, cutter_edges)
            if clipped is not None:
                clipped_lines.append(clipped)


@window.event
def on_draw():
    window.clear()
    glMatrixMode(gl.GL_PROJECTION)
    glLoadIdentity()
    glOrtho(0, window.width, 0, window.height, 0, 1)

    glMatrixMode(gl.GL_MODELVIEW)

    glColor3f(0, 0, 1)
    glBegin(GL_LINES)
    for v in lines_vertices:
        glVertex2f(*v)
    glEnd()

    glColor3f(0, 1, 0)
    glBegin(GL_LINE_STRIP)
    for v in cutter_vertices:
        glVertex2f(*v)

    if len(cutter_vertices) > 2:
        glVertex2f(*cutter_vertices[0])
    glEnd()

    glColor3f(1, 0, 0)
    glBegin(GL_LINES)
    for line in clipped_lines:
        glVertex2f(*line[0])
        glVertex2f(*line[1])
    glEnd()


def clip_line(line, edges):
    t_min = 0
    t_max = 1
    for edge in edges:
        dp = np.dot(line.D, edge.N)
        ddp = np.dot(edge.p1 - line.p1, edge.N)
        t = ddp / dp
        if dp == 0:
            continue
        elif dp > 0:
            t_max = min(t_max, t)
        elif dp < 0:
            t_min = max(t_min, t)

    print("min: ", t_min, "max: ", t_max)

    line = np.array([line.p1 + line.D * t_min, line.p1 + line.D * t_max])
    return line


pyglet.app.run()
