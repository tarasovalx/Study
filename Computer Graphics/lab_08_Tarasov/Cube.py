from Model import *


class Cube(Model):

    @draw_model
    def __draw(self):
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

    def draw(self):
        self.__draw(rot=self.rot)
