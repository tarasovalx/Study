import numpy as np

from math import floor
from pyglet.gl import *
from pyglet.window.key import *

WINDOW_WIDTH, WINDOW_HEIGHT = 800, 800
WIDTH = 100

triangles_count = 6
colors_rot = 0
speed = 0.05
angle = 0

window = pyglet.window.Window(WINDOW_WIDTH, WINDOW_HEIGHT)

triangle_orig_colors = np.array([[1.0, 0.4, 0.2],
                                 [0.2, 1.0, 0.4],
                                 [0.4, 0.2, 1.0]])

triangle_colors = triangle_orig_colors.copy()


def interpolate(a, b, k):
    return a * (1 - k ** 1.2) + b * k ** 1.2


def make_regular_triangle(x, y, a):
    h = a * (3 ** 0.5) / 2
    return np.array([[x - a / 2, y - h / 2],
                     [x, y + h / 2],
                     [x + a / 2, y - h / 2]])


def make_triangles(n=4, h_shift=40):
    width = window.width / 2
    height = window.height / 2
    return [make_regular_triangle(width, height + (h_shift * i) / 10 - 40, WIDTH + h_shift * i) for i in range(n * 2, 0, -1)]


def draw_triangle(t, colored=True):
    colors = triangle_colors if colored else np.zeros((3, 3))
    for i in range(3):
        gl.glColor3f(*colors[i])
        glVertex2f(*t[i])


triangles = make_triangles(triangles_count)


@window.event
def on_draw():
    window.clear()
    glClear(GL_COLOR_BUFFER_BIT)
    glLoadIdentity()
    glBegin(GL_TRIANGLES)

    for i, triangle in enumerate(triangles):
        if i % 2 != 0:
            draw_triangle(triangle)
        else:
            draw_triangle(triangle, False)

    glEnd()


@window.event
def on_mouse_scroll(x, y, scroll_x, scroll_y):
    global triangle_colors, colors_rot, triangle_orig_colors
    colors_rot = (colors_rot + speed * scroll_y) % 3

    for i in range(len(triangle_colors)):
        pos = (i + floor(colors_rot % 3))
        triangle_colors[i] = interpolate(triangle_orig_colors[pos % 3],
                                         triangle_orig_colors[(pos + 1) % 3],
                                         colors_rot % 1)


@window.event
def on_text_motion(motion):
    global triangles_count, triangles

    if motion == MOTION_UP:
        triangles_count = min(40, triangles_count + 1)
        triangles = make_triangles(triangles_count)

    if motion == MOTION_DOWN:
        triangles_count = max(0, triangles_count - 1)
        triangles = make_triangles(triangles_count)


pyglet.app.run()
