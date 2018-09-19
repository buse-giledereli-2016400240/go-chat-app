const socket = io();

var colors = ['#9b009b', '#6189dd', '#66cdaa', '#ff0101', '#0105ff', '#009020', '#ff8300', '#ff0077'];
var random_color = colors[Math.floor(Math.random() * colors.length)];

new Vue({
    el: '#chat-app',
    data: {
        newMsg:'', //holds new messages to be sent to the server
        messageList: [], //running list of chat messages displayed on the screen
        username: '', //our username
        password: '', //our password
        userColor: random_color, //our randomly chosen color
        usersList: [], //running list of usernames displayed on the screen
        joined: false, // True username have been filled in
        entertypechosen: false, //will be true if we click sign up or log in
        loginchosen: false //will be true if we click log in
    },
    created() {
        socket.on("messageSent", (msg) => {
            currDate = new Date();
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
            this.messageList.push({
                username: msg.username,
                userColor: msg.color,
                text: 'joined',
                date: currDate.getHours() + ":" + currDate.getMinutes()
            })
            this.usersList.push({
                username: msg.username,
            })
            var element = document.getElementById('chat-messages');
            element.scrollTop = element.scrollHeight+10; //auto scroll to the bottom
        })
        socket.on("joinedSuccessfully", (msg) => {
            this.joined = true;
            currDate = new Date();
            this.messageList.push({
                username: msg.username,
                userColor: msg.color,
                text: 'joined',
                date: currDate.getHours() + ":" + currDate.getMinutes()
            })
            this.usersList.push({
                username: msg.username,
            })
            var element = document.getElementById('chat-messages');
            element.scrollTop = element.scrollHeight+10; //auto scroll to the bottom
        })
        socket.on("userDisconnected", (uName) => {
            currDate = new Date();
            this.messageList.push({
                username: uName,
                userColor: '',
                text: 'left',
                date: currDate.getHours() + ":" + currDate.getMinutes()
            })
            this.usersList = this.usersList.filter(function(item) { 
                return item.username !== uName
            })
            
            var element = document.getElementById('chat-messages');
            element.scrollTop = element.scrollHeight+10; //auto scroll to the bottom
        })
        socket.on("receiveUserList", (uName) => {
            this.usersList.push({
                username: uName
            })
        })
        socket.on("errorJoining", (errorDesc) => {
            alert(errorDesc)
        })
    },
    methods: {
        sendMessage () {
            if (this.newMsg != '') {
                socket.emit("messageSent", {username: this.username, password: this.password, message: this.newMsg, color: this.userColor});
                this.newMsg = ''; 
            }else {
                alert("You haven't typed a message!");
            }
        },
        join () {
            if(this.username == '' || this.password == '') {
                alert("Username and password required!");
            }else {
                if (this.loginchosen){
                    socket.emit("logInRequest", {username: this.username, password: this.password, message: this.newMsg, color: this.userColor});
                }else{
                    socket.emit("signUpRequest", {username: this.username, password: this.password, message: this.newMsg, color: this.userColor});
                }
            }
        },
    }
})
