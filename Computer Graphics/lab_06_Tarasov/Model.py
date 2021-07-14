from helpers import *

default_animate_keys = array([[0, 0, 0],
                              [2, 0, 0],
                              [2, 0, 2],
                              [-4, 2, 0],

                              [-4, 2, 0],
                              [-2, 0, 1],
                              [0, 0, 1],
                              [1, -1, 1],

                              [1, -1, 1],
                              [1, 0, 0],
                              [1, 0, -2],
                              [0, 2, -4],

                              [0, 2, -4],
                              [2, 0, 2],
                              [0, 0, 2],
                              [0, 0, 0]])


class Model:
    def __init__(self, pos=array([0, 0, 0]), rot=array([0, 0, 0]), scale=array([1, 1, 1]),
                 animate_keys=default_animate_keys):
        self.pos = pos
        self.rot = rot
        self.scale = scale
        self.animate_keys = animate_keys
        self.current_key = 0
        self.acc_t = 0
        self.period = 2

    def animate(self, dt):
        def get_point_ind(i):
            return (self.current_key + i) % len(self.animate_keys)

        def get_point_val(i):
            return self.animate_keys[get_point_ind(i)]

        self.acc_t += dt
        if self.acc_t >= self.period:
            self.acc_t = 0
            self.current_key = get_point_ind(4)

        t = self.acc_t / self.period
        self.pos = get_point_val(0) * (1 - t) ** 3 + \
                   3 * get_point_val(1) * t * (1 - t) ** 2 + \
                   3 * get_point_val(2) * (t ** 2) * (1 - t) + \
                   get_point_val(3) * (t ** 3)

    def save(self):
        return {
            "pos": self.pos,
            "rot": self.rot,
            "scale": self.scale
        }

    @staticmethod
    def load(self, kwargs):
        return Model(self, **kwargs)
