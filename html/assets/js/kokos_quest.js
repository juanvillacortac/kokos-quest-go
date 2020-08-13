let canvas

const WASM_URL = '/assets/wasm/kokos_quest.wasm'
const GZIP = true

window.addEventListener('DOMContentLoaded', async () => {
  const go = new Go()
  const pako = window.pako

  const status = document.getElementById('status')

  status.innerText = 'Fetching'
  const resp = await fetch(GZIP ? WASM_URL + '.gz' : WASM_URL)
  if (!resp.ok) {
    error(await resp.text())
    return
  }
  const buffer = await resp.arrayBuffer()

  let wasm
  if (GZIP) {
    status.innerText = 'Decoding'
    wasm = pako.ungzip(buffer);
    // A fetched response might be decompressed twice on Firefox.
    // See https://bugzilla.mozilla.org/show_bug.cgi?id=610679
    if (wasm[0] === 0x1f && wasm[1] === 0x8b) {
      wasm = pako.ungzip(wasm)
    }
  } else {
    wasm = buffer
  }

  WebAssembly.instantiate(wasm, go.importObject)
    .then(result => {
      document.getElementById('loading').remove()
      status.remove()

      go.run(result.instance)

      canvas = document.getElementsByTagName('canvas')[0]

      SetScreenSize(
        document.__Game__ScreenSize.width,
        document.__Game__ScreenSize.height
      )
    })
    .catch(err => {
      error(err)
    })
})

function error(err) {
  const pre = document.createElement('pre');
  pre.innerText = err
  document.body.appendChild(pre)
  document.getElementById('loading').innerText = 'Error'
  console.error(err)
}

function IsFullscreen() {
  return (document.fullscreenElement && document.fullscreenElement !== null) ||
    (document.webkitFullscreenElement && document.webkitFullscreenElement !== null) ||
    (document.mozFullScreenElement && document.mozFullScreenElement !== null) ||
    (document.msFullscreenElement && document.msFullscreenElement !== null)
}

function SetFullscreen(fullscreen){
  if (canvas == undefined || canvas == null) {
    return
  }

  const enterHandler = () => {
    if (canvas.RequestFullScreen) {
        return canvas.RequestFullScreen()
    } else if (canvas.webkitRequestFullScreen) {
        return canvas.webkitRequestFullScreen()
    } else if(canvas.mozRequestFullScreen) {
        return canvas.mozRequestFullScreen()
    } else if(canvas.msRequestFullscreen) {
        return canvas.msRequestFullscreen()
        alert("This browser doesn't supporter fullscreen")
      alert("This browser doesn't supporter fullscreen")
    }
  }

  const exitHandler = () => {
    if (document.exitFullscreen) {
        return document.exitFullscreen()
    } else if (document.webkitExitFullscreen) {
        return document.webkitExitFullscreen()
    } else if (document.mozCancelFullScreen) {
        return document.mozCancelFullScreen()
    } else if (document.msExitFullscreen) {
        return document.msExitFullscreen()
    }else{
        alert("Exit fullscreen doesn't work")
    }
  }

  if (fullscreen) {
    enterHandler()
  } else {
    exitHandler()
  }
}

const insideInstalledApp = () =>
  window.matchMedia('(display-mode: standalone)').matches || window.navigator.standalone === true

function SetScreenSize(width, height) {
  if (canvas == undefined || canvas == null) {
    return
  }

  if (insideInstalledApp()) {
    if (window.outerWidth) {
      window.resizeTo(
        width + (window.outerWidth - window.innerWidth),
        height + (window.outerHeight - window.innerHeight)
      );
    } else {
      window.resizeTo(500, 500);
      window.resizeTo(
        width + (500 - document.body.offsetWidth),
        height + (500 - document.body.offsetHeight)
      );
    }
  }
}
