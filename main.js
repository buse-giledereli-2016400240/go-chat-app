const socket = io();

var colors = ['#660066', '#3b5998', '#66cdaa', '#ff0101', '#0105ff', '#009020'];
var random_color = colors[Math.floor(Math.random() * colors.length)];

new Vue({
    el: '#chat-app',
    data: {
        newMsg:'', // Holds new messages to be sent to the server
        messageList: [], // A running list of chat messages displayed on the screen
        username: '', // Our username
        userColor: random_color,
        joined: false // True if email and username have been filled in
    },
    created() {
        socket.on("message-sent", (msg) => {
            currDate = new Date();
            console.log(msg);
            this.messageList.push({
                username: msg.username,
                userColor: msg.color,
                text: msg.message,
                date: currDate.getHours() + ":" + currDate.getMinutes()
            })
            var element = document.getElementById('chat-messages');
            element.scrollTop = element.scrollHeight+10; // Auto scroll to the bottom
        })
        socket.on("joined", (msg) => {
            currDate = new Date();
            console.log(msg);
            this.messageList.push({
                username: msg.username,
                userColor: msg.color,
                text: 'joined',
                date: currDate.getHours() + ":" + currDate.getMinutes()
            })
            var element = document.getElementById('chat-messages');
            element.scrollTop = element.scrollHeight+10; // Auto scroll to the bottom
        })
        socket.on("userDisconnected", (uName) => {
            currDate = new Date();
            this.messageList.push({
                username: uName,
                userColor: '',
                text: 'left',
                date: currDate.getHours() + ":" + currDate.getMinutes()
            })
            var element = document.getElementById('chat-messages');
            element.scrollTop = element.scrollHeight+10; // Auto scroll to the bottom
        })
    },
    methods: {
        sendMessage () {
            if (this.newMsg != '') {
                socket.emit("message-sent", {username: this.username, message: this.newMsg, color: this.userColor});
                this.newMsg = ''; // Reset newMsg
            }else {
                alert("You haven't typed a message!");
            }
        },
        join () {
            if(this.username != '') {
                socket.emit("joined", {username: this.username, message: this.newMsg, color: this.userColor});
                this.joined = true;
            }else {
                alert("Username required!");
            }
        },
    }
})
