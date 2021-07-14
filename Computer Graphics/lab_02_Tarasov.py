import numpy as np
from pyglet.gl import *
from pyglet.window.key import *

WINDOW_WIDTH, WINDOW_HEIGHT = 800, 800

window = pyglet.window.Window(WINDOW_WIDTH, WINDOW_HEIGHT)
keyboard = KeyStateHandler()
window.push_handlers(keyboard)

size = 1
rot = [0, 0, 0]
pos = [1, 0, 0]
angle_keys = [X, Y, Z]

polygon_modes = [GL_LINE, GL_FILL]
mode = 0


@window.event
def on_text_motion(motion):
    global rot, mode
    r = [(1 if keyboard[k] else 0) for k in angle_keys]

    if motion == MOTION_UP:
        for i in range(len(rot)):
            rot[i] = ((rot[i] + 5 * r[i]) % 360)

    if motion == MOTION_DOWN:
        for i in range(len(rot)):
            rot[i] = ((rot[i] - 5 * r[i]) % 360)

    if motion == MOTION_NEXT_PAGE:
        mode = not mode


def draw_cube(pos, rot=None):
    glPushMatrix()
    glTranslatef(*pos)
    if rot is not None:
        glRotatef(rot[2], 0, 0, 1)
        glRotatef(rot[1], 0, 1, 0)
        glRotatef(rot[0], 1, 0, 0)
    glScalef(size, size, size)

    # Cторона с удачными цветами
    glRotatef(90, 0, 1, 0)

    glTranslatef(-0.5, -0.5, -0.5)
    glBegin(GL_QUADS)
    glColor3f(0.1, 0.3, 0.6)
    glVertex3f(0, 0, 0)
    glVertex3f(1, 0, 0)
    glVertex3f(1, 0, 1)
    glVertex3f(0, 0, 1)

    glColor3f(0.4, 0.5, 1)
    glVertex3f(0, 1, 0)
    glVertex3f(0, 1, 1)
    glVertex3f(1, 1, 1)
    glVertex3f(1, 1, 0)

    glColor3f(0.4, 0.6, 0.7)
    glVertex3f(0, 0, 0)
    glVertex3f(0, 1, 0)
    glVertex3f(0, 1, 1)
    glVertex3f(0, 0, 1)

    glColor3f(0.6, 0.4, 0.7)
    glVertex3f(1, 0, 0)
    glVertex3f(1, 1, 0)
    glVertex3f(1, 1, 1)
    glVertex3f(1, 0, 1)

    glColor3f(0.6, 0.3, 0.8)
    glVertex3f(1, 0, 0)
    glVertex3f(1, 1, 0)
    glVertex3f(0, 1, 0)
    glVertex3f(0, 0, 0)

    glColor3f(0.3, 0.6, 0.7)
    glVertex3f(1, 0, 1)
    glVertex3f(1, 1, 1)
    glVertex3f(0, 1, 1)
    glVertex3f(0, 0, 1)

    glEnd()
    glPopMatrix()


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


def set_projection():
    glMatrixMode(GL_PROJECTION)
    glLoadIdentity()
    glOrtho(-2, 2, -2, 2, -2, 2)
    glMultMatrixf(My)
    glRotatef(45, 0, 1, 0)
    glMultMatrixf(Mz)
    glRotatef(-90, 0, 1, 0)
    glMultMatrixf(Mx)
    glTranslatef(-1, 0, -1)
    glRotatef(-90, 0, 1, 0)




@window.event
def on_draw():
    window.clear()
    glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT)
    glEnable(GL_DEPTH_TEST)
    glDepthMask(GL_TRUE)

    set_projection()

    glMatrixMode(GL_MODELVIEW)
    glPolygonMode(GL_FRONT_AND_BACK, polygon_modes[mode])

    glLoadIdentity()
    draw_cube(np.array([0, 0, -1.2]))
    draw_cube(np.array([1.2, 0, 0]), rot)


pyglet.app.run()
