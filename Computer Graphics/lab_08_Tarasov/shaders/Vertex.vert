varying vec3 n;
varying vec3 v;

void main()
{   
    v = vec3(gl_ModelViewMatrix * gl_Vertex);
    n = normalize(gl_NormalMatrix * gl_Normal);
    gl_Position = gl_ProjectionMatrix * vec4(v.x, v.y, v.z, 1);
}
