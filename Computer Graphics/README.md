# Computer graphics laboratory works
This repo contains computer graphics laboratory works written on python.


## Launch

All labs uses pyglet and numpy, also lab6 needs numba
to lauch run string below in terminal

```python file_name.py -O```


## Summary 
- Lab 1. OpenGL API basics
- Lab 2. Projection. Three-point projection
- Lab 3. Geometryc modeling. Polygonal meshes. Generation of extrusion mold
- Lab 4. Rasterization of an arbitrary polygon. Scan line by line with a list of active edges
- Lab 5. Clipping algorithms. Kirus-Beck algorithm
- Lab 6. Formation of realistic images. Phong illumination. Texturing.
- Lab 7. OpenGL applications optimisation.
- Lab 8. Shaders. Phong shading.

## Usage
In labs 2, 3, 6, 7, 8 for model rotation you should hold ```x```, ```y``` or ```z```  and press ```ARROW_UP/DOWN``` to rotate arround corresponding axis. ```PAGE_DOWN``` will switch polygon mode (fill or line). To increase/decrease polygons in model press ```ARROW_RIGHT``` or ```ARROW_LEFT```.

In labs 6-8 you can press ```ENTER``` to save scene statement, by pressing ```SPACE``` you will load it.

In lab 4 click ```MOUSE_LEFT``` to enter polugon vertex. To rasterize press ```G```. To clear press ```SPACE```

In lab 5  click ```MOUSE_LEFT``` to enter cutting lines, ```MOUSE_RIGHT``` to enter cutter vertices (you should enter thme clock-wise)


## Demo