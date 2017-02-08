"use strict"

// both enable on Firefox/Chrome
const browser = window.chrome || window.browser

/*
On startup, connect to the "ping_pong" app.
*/
const port = browser.runtime.connectNative("example_golang")

/*
Listen for messages from the app.
*/
port.onMessage.addListener(response => {
  console.log("Received")
  console.log(response)
})

port.onDisconnect.addListener(port => {
  console.log("Disconnected")
  console.log(port)
})

/*
On a click on the browser action, send the app a message.
*/
browser.browserAction.onClicked.addListener(() => {
  console.log("Sending")
  port.postMessage({
    type: "request",
    value: "hi",
  })
})

