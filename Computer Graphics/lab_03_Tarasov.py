from math import *

import numpy as np
import pyglet.window.key
from pyglet.graphics import *
from pyglet.window.key import *

window = pyglet.window.Window(width=800, height=800, resizable=True)
window.set_minimum_size(300, 300)
keyboard = KeyStateHandler()
window.push_handlers(keyboard)

polygon_modes = [GL_LINE, GL_FILL]
angle_keys = [X, Y, Z]
mode = 0

y = 0
N_MIN, N_MAX, M_MIN, M_MAX = 10, 40, 40, 160
rot = [0, 0, 0]
a, b = 2, 0.4
tao = 2 * 2 * pi * a

COLOR_BLUE = (0.004, 0.851, 0.463)
COLOR_GREEN = (0.035, 0.247, 0.608)


def array(*args, **kwargs):
    kwargs.setdefault("dtype", np.float32)
    return np.array(*args, **kwargs)


def draw_model(draw_model_func):
    def decorated(self=None, pos=None, rot=None, scale=None):
        glPushMatrix()
        if pos is not None:
            glTranslatef(*pos)
        if rot is not None:
            glRotatef(rot[2], 0, 0, 1)
            glRotatef(rot[1], 0, 1, 0)
            glRotatef(rot[0], 1, 0, 0)
        if scale is not None:
            glTranslatef(*scale)
        if self is None:
            draw_model_func()
        else:
            draw_model_func(self)
        glPopMatrix()

    return decorated


def curve(t):
    return array([-a * cos(t), b * t, a * sin(t)])


def d_curve(t):
    return array([a * sin(t), b, a * cos(t)])


def dd_curve(t):
    return array([a * cos(t), 0, -a * sin(t)])


def create_freinet_basis(t, d_curve=d_curve, dd_curve=dd_curve):
    d = d_curve(t) / np.sqrt(np.sum(d_curve(t) ** 2))
    dd = dd_curve(t) / np.sqrt(np.sum(dd_curve(t) ** 2))

    b = np.cross(d, dd) / np.sqrt(np.sum(np.cross(d, dd) ** 2))
    c = np.cross(b, d) / np.sqrt(np.sum(np.cross(b, d) ** 2))
    return np.stack([c, b, d]).T


def create_circle(frag=N, a=b):
    phis = np.linspace(0, 2 * pi, num=frag)
    return np.stack([a * np.cos(phis), a * np.sin(phis), np.repeat(0, frag)]).T


def create_spiral(x_frag=N, y_frag=M):
    y_frag *= 2
    circle = create_circle(x_frag).T
    vertices = np.ndarray(shape=(y_frag, x_frag, 3), dtype=np.float32)
    ts = np.linspace(-tao, tao, num=y_frag)
    for i, t in np.ndenumerate(ts):
        vertices[i] = (create_freinet_basis(t) @ circle).T + curve(t)

    return vertices


@draw_model
def draw_cube():
    glTranslatef(*(-1) * array([0.5, 0.5, 0.5]))
    glBegin(GL_QUADS)

    glColor3f(0.2, 0.5, 1)
    glVertex3f(*[0, 0, 0])
    glVertex3f(*[1, 0, 0])
    glVertex3f(*[1, 0, 1])
    glVertex3f(*[0, 0, 1])

    glColor3f(0.5, 0.5, 1)
    glVertex3f(*[0, 1, 0])
    glVertex3f(*[0, 1, 1])
    glVertex3f(*[1, 1, 1])
    glVertex3f(*[1, 1, 0])

    glColor3f(0.5, 0.7, 0.3)
    glVertex3f(*[0, 0, 0])
    glVertex3f(*[0, 1, 0])
    glVertex3f(*[0, 1, 1])
    glVertex3f(*[0, 0, 1])

    glColor3f(0.5, 0.7, 0.3)
    glVertex3f(*[1, 0, 0])
    glVertex3f(*[1, 1, 0])
    glVertex3f(*[1, 1, 1])
    glVertex3f(*[1, 0, 1])

    glColor3f(0.9, 0.5, 1)
    glVertex3f(*[1, 0, 0])
    glVertex3f(*[1, 1, 0])
    glVertex3f(*[0, 1, 0])
    glVertex3f(*[0, 0, 0])

    glColor3f(0.5, 0.9, 1)
    glVertex3f(*[1, 0, 1])
    glVertex3f(*[1, 1, 1])
    glVertex3f(*[0, 1, 1])
    glVertex3f(*[0, 0, 1])
    glEnd()


Mx = (gl.GLfloat * 16)(*[1, 0, 0, -0.2,
                         0, 1, 0, 0,
                         0, 0, 1, 0,
                         0, 0, 0, 1])

Mz = (gl.GLfloat * 16)(*[1, 0, 0, 0.2,
                         0, 1, 0, 0,
                         0, 0, 1, 0,
                         0, 0, 0, 1])

My = (gl.GLfloat * 16)(*[1, 0, 0, 0,
                         0, 1, 0, 0.12,
                         0, 0, 1, 0,
                         0, 0, 0, 1])


def set_projection(width, height):
    glMatrixMode(GL_PROJECTION)
    glLoadIdentity()
    aspect_ratio = width / height
    x, y, z = aspect_ratio * 8, 8, 20
    glOrtho(-x / 2, x / 2, -y / 2 + 2, y / 2 + 2, -z / 2, z / 2)
    glMultMatrixf(My)
    glRotatef(45, 0, 1, 0)
    glMultMatrixf(Mz)
    glRotatef(-90, 0, 1, 0)
    glMultMatrixf(Mx)
    glTranslatef(-1, 0, -1)


def draw_cords(cords_size=100):
    glBegin(GL_LINES)
    glColor3f(1, 0, 0)
    glVertex3f(-cords_size, 0, 0)
    glVertex3f(cords_size, 0, 0)
    glColor3f(0, 1, 0)
    glVertex3f(0, -cords_size, 0)
    glVertex3f(0, cords_size, 0)
    glColor3f(0, 0, 1)
    glVertex3f(0, 0, -cords_size)
    glVertex3f(0, 0, cords_size)
    glEnd()


@window.event
def on_resize(width, height):
    glViewport(0, 0, width, height)
    return pyglet.event.EVENT_HANDLED


def reorder_vertices(vertices):
    vertices_num = (vertices.shape[0] - 1) * (vertices.shape[1] - 1) * 3
    reordered_vertices = np.ndarray(shape=(vertices_num, 3), dtype=np.float32)
    colors = np.ndarray(shape=(vertices_num, 3), dtype=np.float32)

    c_grad = np.linspace(COLOR_BLUE, COLOR_GREEN, (vertices.shape[1] // 2))
    color_c = np.concatenate([c_grad[::-1], c_grad])
    cnt = 0
    for i in range(0, vertices.shape[0] - 1):
        for j in range(0, vertices.shape[1] - 1, 2):
            reordered_vertices[cnt] = vertices[i][j]
            reordered_vertices[cnt + 1] = vertices[i + 1][j]
            reordered_vertices[cnt + 2] = vertices[i][j + 1]
            reordered_vertices[cnt + 3] = vertices[i + 1][j + 1]

            colors[cnt] = color_c[j]
            colors[cnt + 1] = color_c[j]
            colors[cnt + 2] = color_c[j + 1]
            colors[cnt + 3] = color_c[j + 1]
            cnt += 4
    return cnt, reordered_vertices[:cnt], colors[:cnt]


class Spiral:
    def __init__(self, n, m):
        self.N = max(min(n, N_MAX), N_MIN)
        self.M = max(min(m, M_MAX), M_MIN)
        _model = create_spiral(self.N, self.M)
        vertices_num, self.vertices, self.colors = reorder_vertices(_model)
        self.v_list = pyglet.graphics.vertex_list(vertices_num,
                                                  ('v3f/static', self.vertices.flatten()),
                                                  ('c3f/static', self.colors.flatten()))

    @draw_model
    def draw(self):
        glBegin(GL_TRIANGLE_STRIP)
        for i in range(len(self.vertices)):
            glColor3f(*self.colors[i])
            glVertex3f(*self.vertices[i])
        glEnd()

    @draw_model
    def draw_va(self):
        self.v_list.draw(pyglet.gl.GL_TRIANGLE_STRIP)


model = Spiral(10, 40)

glEnable(GL_DEPTH_TEST)
glPolygonMode(GL_FRONT_AND_BACK, polygon_modes[mode])


@window.event
def on_draw():
    window.clear()
    glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT)

    set_projection(window.width, window.height)
    glMatrixMode(GL_MODELVIEW)
    glLoadIdentity()

    model.draw_va(rot=rot)

    draw_cube(pos=array([1.2, 0, 0]))
    draw_cube(pos=array([0, 0, -1.2]))
    draw_cords()
    return pyglet.event.EVENT_HANDLED


@window.event
def on_text_motion(motion):
    global rot, mode, model
    r = [(1 if keyboard[k] else 0) for k in angle_keys]

    if motion == MOTION_UP:
        for i in range(len(rot)):
            rot[i] = ((rot[i] + 5 * r[i]) % 360)

    if motion == MOTION_DOWN:
        for i in range(len(rot)):
            rot[i] = ((rot[i] - 5 * r[i]) % 360)

    if motion == MOTION_RIGHT:
        model = Spiral(model.N + 2, model.M + 8)

    if motion == MOTION_LEFT:
        model = Spiral(model.N - 2, model.M - 8)

    if motion == MOTION_NEXT_PAGE:
        mode = not mode
        glPolygonMode(GL_FRONT_AND_BACK, polygon_modes[mode])


def update(dt):
    global rot
    rot += array([0, 20 * dt, 0])


pyglet.clock.schedule_interval(update, 0.005)
pyglet.app.run()
