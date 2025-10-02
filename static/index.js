const name = prompt("Enter your name:") || "Anonymous";
const protocol = location.protocol === 'https:' ? 'wss' : 'ws';
const ws = new WebSocket(`${protocol}://${location.host}/ws`);
const msgInput = document.getElementById("msg");
const sendBtn = document.getElementById("send");
const messagesList = document.getElementById("messages");

// Markdown sanitizer (only allow **bold**, *italic*, `code`)
function renderMarkdown(text) {
    let safe = text
        .replace(/</g, "&lt;") // escape HTML
        .replace(/>/g, "&gt;");
    safe = safe.replace(/\*\*(.*?)\*\*/g, "<strong>$1</strong>"); // bold
    safe = safe.replace(/\*(.*?)\*/g, "<em>$1</em>");             // italic
    safe = safe.replace(/`(.*?)`/g, "<code>$1</code>");           // inline code
    safe = safe.replace(/\n/g, "<br>");                           // new line
    return safe;
}

ws.onopen = () => {
    ws.send(name); // Send name as first message
};

ws.onclose = () => {
    alert("Connection is closed.")
}

ws.onmessage = (event) => {
    const { name, message, timestamp, system } = JSON.parse(event.data);
    const item = document.createElement("li");

    if (system) {
        item.className = "system";
        item.textContent = `${name} ${message}`;
    } else {
        const safeMessage = renderMarkdown(message);
        item.innerHTML = `<span style="color:#1976d2;font-weight:bold;">${name}</span>: ${safeMessage}`;
    }

    messagesList.appendChild(item);
    messagesList.scrollTop = messagesList.scrollHeight;
};

sendBtn.onclick = () => {
    const message = msgInput.value.trim();
    if (message) {
        ws.send(JSON.stringify({ message }));
        msgInput.value = "";
    }
};

msgInput.addEventListener("keydown", (e) => {
    if (e.key === "Enter" && !e.shiftKey) {
        e.preventDefault();
        sendBtn.click();
    }
});