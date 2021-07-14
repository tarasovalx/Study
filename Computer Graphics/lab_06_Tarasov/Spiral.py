import pyglet.window.key
from math import pi, cos, sin

from Model import *

N_MIN, N_MAX, M_MIN, M_MAX = 10, 40, 40, 160
a, b = 2, 0.4
tao = 2 * 2 * pi * a

COLOR_BLUE = (0.004, 0.851, 0.463)
COLOR_GREEN = (0.035, 0.247, 0.608)

N, M = 100, 10


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


def create_spartial(x_frag=N, y_frag=M):
    y_frag *= 2
    circle = create_circle(x_frag).T
    vertecies = np.ndarray(shape=(y_frag, x_frag, 3), dtype=np.float32)
    normals = np.ndarray(shape=(y_frag, x_frag, 3), dtype=np.float32)
    ts = np.linspace(-tao, tao, num=y_frag)
    for i, t in enumerate(ts):
        vertecies[i] = (create_freinet_basis(t) @ circle).T + curve(t)
        normals[i] = vertecies[i] - curve(t)
    return vertecies, normals


def reorder_vertices(vertices, normals, n, m):
    vertices_num = (vertices.shape[0] - 1) * (vertices.shape[1] - 1) * 3
    vertices_arr = np.ndarray(shape=(vertices_num, 3), dtype=np.float32)
    colors_arr = np.ndarray(shape=(vertices_num, 3), dtype=np.float32)
    normals_arr = np.ndarray(shape=(vertices_num, 3), dtype=np.float32)
    tex_coords_arr = np.ndarray(shape=(vertices_num, 2), dtype=np.float32)

    c_grad = np.linspace(COLOR_BLUE, COLOR_GREEN, (vertices.shape[1] // 2))
    color_c = np.concatenate([c_grad[::-1], c_grad])
    cnt = 0
    for i in range(0, vertices.shape[0] - 1):
        for j in range(0, vertices.shape[1] - 1, 2):
            vertices_arr[cnt] = vertices[i][j]
            vertices_arr[cnt + 1] = vertices[i + 1][j]
            vertices_arr[cnt + 2] = vertices[i][j + 1]
            vertices_arr[cnt + 3] = vertices[i + 1][j + 1]

            normals_arr[cnt] = normals[i][j]
            normals_arr[cnt + 1] = normals[i + 1][j]
            normals_arr[cnt + 2] = normals[i][j + 1]
            normals_arr[cnt + 3] = normals[i + 1][j + 1]

            colors_arr[cnt] = color_c[j]
            colors_arr[cnt + 1] = color_c[j]
            colors_arr[cnt + 2] = color_c[j + 1]
            colors_arr[cnt + 3] = color_c[j + 1]

            tex_coords_arr[cnt] = [i / n * 4, j / m * 4]
            tex_coords_arr[cnt + 1] = [(i + 1) / n * 4, j / m * 4]
            tex_coords_arr[cnt + 2] = [i / n * 4, (j + 1) / m * 4]
            tex_coords_arr[cnt + 3] = [(i + 1) / n * 4, (j + 1) / m * 4]

            cnt += 4
    return cnt, vertices_arr[:cnt], colors_arr[:cnt], normals_arr[:cnt], tex_coords_arr[:cnt]


class Spiral(Model):
    def __init__(self, n, m, **kwargs):
        super().__init__(**kwargs)
        self.N = max(min(n, N_MAX), N_MIN)
        self.M = max(min(m, M_MAX), M_MIN)
        self.__gen_vertices()

    def __gen_vertices(self):
        model, normals = create_spartial(self.N, self.M)
        vertices_num, self.vertices, self.colors, self.normals, self.tex_coords = reorder_vertices(model, normals, self.N, self.M)
        self.v_list = pyglet.graphics.vertex_list(vertices_num,
                                                  ('v3f/static', self.vertices.flatten()),
                                                  ('c3f/static', self.colors.flatten()),
                                                  ('n3f/static', self.normals.flatten()),
                                                  ('t2f/static', self.tex_coords.flatten()))

    def update_frag(self, n=None, m=None):
        if n is not None:
            self.N = max(min(n, N_MAX), N_MIN)
        if m is not None:
            self.M = max(min(m, M_MAX), M_MIN)

        self.__gen_vertices()

    @draw_model
    def __draw(self):
        glBegin(GL_TRIANGLE_STRIP)
        for i in range(len(self.vertices)):
            glColor3f(*self.colors[i])
            glNormal3f(*self.normals[i])
            glTexCoord2d(*self.tex_coords[i])
            glVertex3f(*self.vertices[i])
        glEnd()

    def draw(self):
        self.__draw_va(rot=self.rot, scale=self.scale, pos=self.pos)
#       self.__draw(rot=self.rot, scale=self.scale, pos=self.pos)

    @draw_model
    def __draw_va(self):
        self.v_list.draw(GL_TRIANGLE_STRIP)

    def save(self):
        return {**{"n": self.N, "m": self.M}, **super().save()}
