const ASSETS = [
  '/index.html',
  '/',

  '/assets/css/style.css',

  '/assets/wasm/kokos_quest.wasm.gz',

  '/assets/js/pako@1.0.11.min.js',
  '/assets/js/kokos_quest.js',
  '/assets/js/wasm_exec.js',

  '/assets/metadata/manifest.json',
  '/assets/metadata/apple-touch-icon.png',
  '/assets/metadata/favicon-16x16.png',
  '/assets/metadata/favicon-32x32.png',
  '/assets/metadata/favicon-192x192.png',
  '/assets/metadata/favicon-512x512.png',
  '/assets/metadata/favicon.ico',
]

const CACHE_NAME = 'kokos_quest-v1'

self.addEventListener('install', event => {
  console.info('[Service Worker] Installing...')
  event.waitUntil(precache())
})

function precache() {
    caches
      .open(CACHE_NAME)
      .then(cache => {
        console.info('[Service Worker] Caching...')
        return cache.addAll(ASSETS)
      })
      .catch(err => console.info('[Service Worker] Error caching:', err))
}

self.addEventListener('activate', function(event) {
  console.info('[Service Worker] Activated!')
  event.waitUntil(
    caches.keys().then(function(keyList) {
      return Promise.all(keyList.map(function(key) {
        if (key !== CACHE_NAME) {
          console.info('[Service Worker] Removing old cache...', key)
          return caches.delete(key)
        }
      }))
    })
  )
})

self.addEventListener('fetch', event => {
  console.info('[Service Worker] Fetching', event.request.url)
  event.respondWith(fromCache(event.request))
  event.waitUntil(update(event.request))
})

async function fromCache(request) {
  return caches.open(CACHE_NAME).then(async cache => {
    return cache.match(request).then(response => {
      return response || fetch(request)
    }).catch(function (error) {
      console.info('[Service Worker] Error fetching', request.url, error)
    })
  })
}

async function update(request) {
  return caches.open(CACHE_NAME).then(async cache => {
    return fetch(request).then(response => {
      console.info("[Service Worker] Updating cache...", request.url)
      return cache.put(request, response)
    })
  })
}
