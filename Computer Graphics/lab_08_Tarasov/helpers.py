import numpy as np
from pyglet.gl import *


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
            glScalef(*scale)
        if self is not None:
            draw_model_func(self)
        else:
            draw_model_func()
        glPopMatrix()

    return decorated
