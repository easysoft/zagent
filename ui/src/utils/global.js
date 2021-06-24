import Vue from 'vue'

const EventBus = new Vue()

export default {
  EventBus,

  ws: {},
  wsEventName: 'wsEvent',
  wsDefaultRoom: 'square',
  setWs: function (conn) {
    this.ws = conn
  }
}
