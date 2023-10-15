<template>
  <div>
    <form :action="sendMessage" @click.prevent="onSubmit">
      <img
        class="imgShow"
        src="https://www.ideamotive.co/hubfs/Go_2.png"
        alt=""
      />
      <input v-model="message" type="text" />
      <input type="submit" value="Send" @click="sendMessage" />
    </form>
  </div>
  <div v-if="showMessage">
    <h3>Message in a WebSocket</h3>
    <p>{{ rcvMessage }}</p>
    <button @click="showMessage = !showMessage">Dismiss</button>
  </div>
</template>

<script>


export default {
  name: "App",
  data () {
    return {
      message: "",
      socket:null,
      rcvMessage: "",
      showMessage: false
    }
  },
  mounted() {
    this.socket = new WebSocket("ws://localhost:9100/socket")
    this.socket.onmessage = (msg) => {
      this.acceptMessage(msg)
    }
    
  },

  methods: {
    sendMessage() {
      let msg = {
        "greeting": this.message
      }
      this.socket.send(JSON.stringify(msg))
    },
    acceptMessage(msg) {
      this.rcvMessage = msg.data
      this.showMessage = true
    }
  }
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

.imgShow {
  width: 70rem;
}
</style>
