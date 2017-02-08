"use strict"
/*
On startup, connect to the "ping_pong" app.
*/

// both enable on Firefox/Chrome
const browser = window.chrome || window.browser

const onResponse = response => {
  console.log("Received")
  console.log(response)
}

/*
On a click on the browser action, send the app a message.
*/
browser.browserAction.onClicked.addListener(() => {
  console.log("Sending")
  browser.runtime.sendNativeMessage("example_golang", {
    type: "request",
    value: "hi",
  }, onResponse)
})
