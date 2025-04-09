import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import { VitePWA } from 'vite-plugin-pwa';

export default defineConfig({
  plugins: [
    react(),
    VitePWA({
      registerType: 'autoUpdate',
      inject: {
        injectManifest: true,
      },
      workbox: {
        runtimeCaching: [
          {
            urlPattern: ({ request }) => request.destination === 'document' || request.destination === 'script' || request.destination === 'style',
            handler: 'CacheFirst',
            options: {
              cacheName: 'react-assets-cache',
            },
          },
          {
            // Explicitly match image file extensions (jpg, png, jpeg, gif, svg, webp)
            urlPattern: ({ request }) =>
              request.url.match(/\.(jpg|jpeg|png|gif|svg|webp)$/),
            handler: 'CacheFirst',
            options: {
              cacheName: 'image-cache',
              expiration: {
                maxEntries: 50, // Limit the number of cached images
                maxAgeSeconds: 60 * 60 * 24 * 30, // Cache for 30 days
              },
            },
          },
          {
            urlPattern: ({ request }) => request.pathname.startsWith('/api/'),
            handler: 'NetworkFirst',
            options: {
              cacheName: 'api-cache',
              networkTimeoutSeconds: 3,
            },
          },
          // Font caching rule
          {
            urlPattern: /^https:\/\/pwa\.nathan-strobl\.org\/fonts\/.*\.(woff2?|ttf|otf|eot)$/,
            handler: 'CacheFirst',
            options: {
              cacheName: 'font-cache',
              expiration: {
                maxEntries: 50, // Limit the number of cached font files
                maxAgeSeconds: 60 * 60 * 24 * 30, // Cache for 30 days
              },
            },
          },
        ],
      },
    }),
  ],
});
