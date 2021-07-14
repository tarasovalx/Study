import pyglet
import pyglet.window.key
from pyglet import image
from pyglet.graphics import *
from pyglet.window.key import *
from pyglet.gl import *
import numpy as np
import pickle
from math import *

from helpers import *
from Cube import *
from Model import *
from Spiral import *

window = pyglet.window.Window(width=800, height=800, resizable=True)
window.set_minimum_size(300, 300)
keyboard = KeyStateHandler()

polygon_modes = [GL_LINE, GL_FILL]
angle_keys = [X, Y, Z]

mode = 0
y = 0


class Scene:
    def __init__(self, models=[], lightings=[]):
        self.models = models
        self.config_light()
        self.config_textures()

    def draw(self):
        for model in self.models:
            model.draw()

    def rotate(self, rot):
        for model in self.models:
            model.rot = model.rot + rot

    def animate(self, dt):
        for model in self.models:
            model.animate(dt)

    @staticmethod
    def config_light():
        glEnable(GL_LIGHTING)

        glLightModelfv(GL_LIGHT_MODEL_AMBIENT, (GLfloat * 4)(0, 0, 0, 1))
        glLightModelf(GL_LIGHT_MODEL_LOCAL_VIEWER, GL_TRUE)
        glLightModelfv(GL_LIGHT_MODEL_LOCAL_VIEWER, (GLfloat * 4)(-1, 0, -1, 1))

        glLightfv(GL_LIGHT0, GL_POSITION, (GLfloat * 4)(4, 0, 4, 1))
        glLightfv(GL_LIGHT0, GL_DIFFUSE, (GLfloat * 4)(0.5, 0.3, 0.5, 1))
        glLightfv(GL_LIGHT0, GL_SPOT_DIRECTION, (GLfloat * 3)(-1, 0, -1))
        glLightf(GL_LIGHT0, GL_LINEAR_ATTENUATION, 0.05)
        glLightf(GL_LIGHT0, GL_QUADRATIC_ATTENUATION, 0.02)
        glEnable(GL_LIGHT0)

        glLightfv(GL_LIGHT1, GL_POSITION, (GLfloat * 4)(-8, 8, 0, 1))
        glLightfv(GL_LIGHT1, GL_SPECULAR, (GLfloat * 4)(0, 0, 0.8, 1))
        glLightfv(GL_LIGHT1, GL_DIFFUSE, (GLfloat * 4)(0, 0, 1, 1))
        glLightfv(GL_LIGHT1, GL_SPOT_DIRECTION, (GLfloat * 3)(2.5, -2, 0))
        glLightfv(GL_LIGHT1, GL_SPOT_CUTOFF, (GLfloat)(25.0))

        glEnable(GL_LIGHT1)

        glLightfv(GL_LIGHT2, GL_POSITION, (GLfloat * 4)(8, -8, 0, 1))
        glLightfv(GL_LIGHT2, GL_SPECULAR, (GLfloat * 4)(0.5, 0, 0, 1))
        glLightfv(GL_LIGHT2, GL_DIFFUSE, (GLfloat * 4)(1, 0, 0, 1))
        glLightfv(GL_LIGHT2, GL_SPOT_DIRECTION, (GLfloat * 3)(-2.5, 2, 0))
        glLightfv(GL_LIGHT2, GL_SPOT_CUTOFF, (GLfloat)(25.0))
        glEnable(GL_LIGHT2)

        glMaterialfv(GL_FRONT, GL_DIFFUSE, (GLfloat * 3)(1, 0.7, 1))
        glMaterialfv(GL_FRONT, GL_SPECULAR, (GLfloat * 3)(1, 0, 1))
        glMaterialf(GL_FRONT, GL_SHININESS, 16)
        glMaterialfv(GL_FRONT, GL_EMISSION, (GLfloat * 4)(0.01, 0, 0.01, 1))

    @staticmethod
    def config_textures():
        pic = image.load("texture.bmp")
        texture = pic.get_image_data()
        glEnable(GL_TEXTURE_2D)
        textures_ids = (pyglet.gl.GLuint * 1)()
        glGenTextures(1, textures_ids)
        tex_id = textures_ids[0]
        glBindTexture(GL_TEXTURE_2D, tex_id)

        glTexParameterf(GL_TEXTURE_2D, GL_TEXTURE_WRAP_S, GL_REPEAT)
        glTexParameterf(GL_TEXTURE_2D, GL_TEXTURE_WRAP_T, GL_REPEAT)
        glTexParameterf(GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_NEAREST)
        glTexParameterf(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_NEAREST)

        color_format = 'RGBA'
        pitch = texture.width * len(color_format)
        glTexImage2D(GL_TEXTURE_2D, 0, GL_RGB, texture.width, texture.height,
                     0, GL_RGBA, GL_UNSIGNED_BYTE, pic.get_data(color_format, pitch))

    def save(self):
        return {"models": [(type(model), model.save()) for model in self.models]}

    @staticmethod
    def load(data):
        models = [t(**kwargs) for t, kwargs in data["models"]]
        return Scene(models)


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
    glOrtho(-x / 2, x / 2, -y / 2, y / 2, -z / 2, z / 2)
    glMultMatrixf(My)
    glRotatef(45, 0, 1, 0)
    glMultMatrixf(Mz)
    glRotatef(-90, 0, 1, 0)
    glMultMatrixf(Mx)
    glTranslatef(-1, 0, -1)


@window.event
def on_resize(width, height):
    glViewport(0, 0, width, height)
    return pyglet.event.EVENT_HANDLED


glEnable(GL_DEPTH_TEST)
glPolygonMode(GL_FRONT_AND_BACK, polygon_modes[mode])
glEnable(GL_NORMALIZE)

model = Spiral(10, 40, pos=array([0, 0, 0]), rot=array([90, -45, 0]))
scene = Scene(models=[model])


@window.event
def on_draw():
    window.clear()
    glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT)
    set_projection(window.width, window.height)
    glMatrixMode(GL_MODELVIEW)
    glLoadIdentity()
    scene.draw()

    return pyglet.event.EVENT_HANDLED


@window.event
def on_text_motion(motion):
    global mode, scene
    r = array([(1 if keyboard[k] else 0) for k in angle_keys])

    if motion == MOTION_UP:
        scene.rotate(5 * r)

    if motion == MOTION_DOWN:
        scene.rotate(-5 * r)

    if motion == MOTION_RIGHT:
        N = scene.models[0].N
        M = scene.models[0].M
        scene.models[0].update_frag(N + 2, M + 8)

    if motion == MOTION_LEFT:
        N = scene.models[0].N
        M = scene.models[0].M
        scene.models[0].update_frag(N - 2, M - 8)

    if motion == MOTION_NEXT_PAGE:
        mode = not mode
        glPolygonMode(GL_FRONT_AND_BACK, polygon_modes[mode])


@window.event
def on_key_press(symbol, modifiers):
    global scene
    if symbol == ENTER:
        data = scene.save()
        with open('data.pickle', 'wb') as file:
            pickle.dump(data, file)

    if symbol == SPACE:
        with open('data.pickle', 'rb') as file:
            scene = Scene.load(pickle.load(file))
            pyglet.clock.schedule(scene.animate)

    return pyglet.event.EVENT_HANDLED


window.push_handlers(keyboard)
pyglet.clock.schedule(scene.animate)
pyglet.app.run()
