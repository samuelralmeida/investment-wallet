document.addEventListener('htmx:afterRequest', function(evt) {
    // console.log(evt)

    if (evt.detail.failed) {
        const obj = JSON.parse(evt.detail.xhr.response)
        alert(obj.message)
        return
    }

    if (evt.detail.pathInfo.requestPath.startsWith("/elements")) {
        return
    }

    const element = document.getElementById(evt.target.id)
    element.reset()
});
