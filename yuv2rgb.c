#include <vpx/vpx_image.h>
#include <tox/toxav.h>

void
i420_to_rgb(int width, int height, const uint8_t *y, const uint8_t *u, const uint8_t *v,
            int ystride, int ustride, int vstride, unsigned char *out)
{
    const int w = width;
    const int w2 = w / 2;
    const int pstride = w * 3;
    const int h = height;
    const int h2 = h / 2;

    const int strideY = ystride;
    const int strideU = ustride;
    const int strideV = vstride;
    int posy, posx;

    for (posy = 0; posy < h2; posy++) {
        unsigned char *dst = out + pstride * (posy * 2);
        unsigned char *dst2 = out + pstride * (posy * 2 + 1);
        const unsigned char *srcY = y + strideY * posy * 2;
        const unsigned char *srcY2 = y + strideY * (posy * 2 + 1);
        const unsigned char *srcU = u + strideU * posy;
        const unsigned char *srcV = v + strideV * posy;

        for (posx = 0; posx < w2; posx++) {
            unsigned char Y, U, V;
            short R, G, B;
            short iR, iG, iB;

            U = *(srcU++);
            V = *(srcV++);
            iR = (351 * (V - 128)) / 256;
            iG = - (179 * (V - 128)) / 256 - (86 * (U - 128)) / 256;
            iB = (444 * (U - 128)) / 256;

            Y = *(srcY++);
            R = Y + iR ;
            G = Y + iG ;
            B = Y + iB ;
            R = (R < 0 ? 0 : (R > 255 ? 255 : R));
            G = (G < 0 ? 0 : (G > 255 ? 255 : G));
            B = (B < 0 ? 0 : (B > 255 ? 255 : B));
            *(dst++) = R;
            *(dst++) = G;
            *(dst++) = B;

            Y = *(srcY2++);
            R = Y + iR ;
            G = Y + iG ;
            B = Y + iB ;
            R = (R < 0 ? 0 : (R > 255 ? 255 : R));
            G = (G < 0 ? 0 : (G > 255 ? 255 : G));
            B = (B < 0 ? 0 : (B > 255 ? 255 : B));
            *(dst2++) = R;
            *(dst2++) = G;
            *(dst2++) = B;

            Y = *(srcY++) ;
            R = Y + iR ;
            G = Y + iG ;
            B = Y + iB ;
            R = (R < 0 ? 0 : (R > 255 ? 255 : R));
            G = (G < 0 ? 0 : (G > 255 ? 255 : G));
            B = (B < 0 ? 0 : (B > 255 ? 255 : B));
            *(dst++) = R;
            *(dst++) = G;
            *(dst++) = B;

            Y = *(srcY2++);
            R = Y + iR ;
            G = Y + iG ;
            B = Y + iB ;
            R = (R < 0 ? 0 : (R > 255 ? 255 : R));
            G = (G < 0 ? 0 : (G > 255 ? 255 : G));
            B = (B < 0 ? 0 : (B > 255 ? 255 : B));
            *(dst2++) = R;
            *(dst2++) = G;
            *(dst2++) = B;
        }
    }
}

void rgb_to_i420(unsigned char* rgb, vpx_image_t *img)
{
    int upos = 0;
    int vpos = 0;
    int x = 0, i = 0;
    int line = 0;

    for (line = 0; line < img->d_h; ++line) {
        if (!(line % 2)) {
            for (x = 0; x < img->d_w; x += 2) {
                uint8_t r = rgb[3 * i];
                uint8_t g = rgb[3 * i + 1];
                uint8_t b = rgb[3 * i + 2];

                img->planes[VPX_PLANE_Y][i++] = ((66*r + 129*g + 25*b) >> 8) + 16;
                img->planes[VPX_PLANE_U][upos++] = ((-38*r + -74*g + 112*b) >> 8) + 128;
                img->planes[VPX_PLANE_V][vpos++] = ((112*r + -94*g + -18*b) >> 8) + 128;

                r = rgb[3 * i];
                g = rgb[3 * i + 1];
                b = rgb[3 * i + 2];

                img->planes[VPX_PLANE_Y][i++] = ((66*r + 129*g + 25*b) >> 8) + 16;
            }
        } else {
            for (x = 0; x < img->d_w; x += 1) {
                uint8_t r = rgb[3 * i];
                uint8_t g = rgb[3 * i + 1];
                uint8_t b = rgb[3 * i + 2];

                img->planes[VPX_PLANE_Y][i++] = ((66*r + 129*g + 25*b) >> 8) + 16;
            }
        }
    }
}
