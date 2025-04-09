const CACHE_NAME = 'walmarks-cache';
const cacheUrls = [
    '/',
    '/cart',
    '/browse',
    // as many files as you need to cache can go here
];

self.addEventListener('install', event => event.waitUntil(caches.open(CACHE_NAME).then(cache => cache.addAll(cacheUrls))));

self.addEventListener('fetch', event => {
    event.respondWith(
        caches.match(event.request).then(cachedResponse => { 
            if(cachedResponse) return cachedResponse.clone();

            return fetch(event.request).then(networkResponse => {
                const clonedResponse = networkResponse.clone();

                if(event.request.url.startsWith(self.location.origin) && networkResponse && networkResponse.status === 200) {
                    caches.open(CACHE_NAME).then(cache => {
                        cache.put(event.request, clonedResponse);
                    });
                }

                return networkResponse;
            }); 
        })
    );
});