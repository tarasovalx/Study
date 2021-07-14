varying vec3 n;
varying vec3 v;

void main(void)
{
    vec4 result = vec4(0.0);
    vec4 Iamb, Idiff, Ispec; 
    vec3 lightDir, viewDir, reflectDir;
    float diffuseAngle, specularAngle;
    
    for (int i = 0; i < gl_MaxLights; ++i)
    {
        if (gl_LightSource[i].position.w != 0.0)
        {
            lightDir = normalize(gl_LightSource[i].position.xyz - v);
        }
        else
        {
            lightDir = normalize(gl_LightSource[i].position.xyz);
        }
        viewDir = normalize(-v);
        reflectDir = normalize(-reflect(lightDir, n));

        Iamb = gl_FrontLightProduct[i].ambient;

        diffuseAngle = max(dot(n, lightDir), 0.0);
        Idiff = gl_FrontLightProduct[i].diffuse * diffuseAngle;
        Idiff = clamp(Idiff, 0.0, 1.0);

        specularAngle = max(dot(reflectDir, viewDir), 0.0);
        Ispec = gl_FrontLightProduct[i].specular * pow(specularAngle, gl_FrontMaterial.shininess);
        Ispec = clamp(Ispec, 0.0, 1.0);

        result += Iamb + Idiff + Ispec;
    }

    gl_FragColor = gl_FrontLightModelProduct.sceneColor + result;
}