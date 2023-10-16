import { defineStore } from 'pinia';

import { fetchWrapper }  from '../helpers/fetch-wrapper.js';
import { useUsersStore } from './user.store.js';

const chatURL = `${import.meta.env.VITE_API_URL}/chat`;
const wsURL = `${import.meta.env.VITE_WS_URL}/chat/socket`;

export const useChatStore = defineStore({
  id: 'chat',
  state: () => ({
    messages: [
      {
        sender: "ai",
        content: "Hello there... I'm Asai, How can I help you?"
      }
    ],
    socket: null
  }),
  actions: {
    async connectWebSocket() {
      this.socket = new WebSocket(wsURL);

      this.socket.addEventListener('open', (event) => {
        console.log('WebSocket connected', event);
      });

      this.socket.addEventListener('message', (event) => {
        console.log('Received message');
        var msg = {
          sender: "ai",
          content: event.data
        }
        this.messages = [...this.messages, msg];
      });

      this.socket.addEventListener('close', (event) => {
        console.log('WebSocket closed', event);
      });
    },
    async loadHistory() {
      console.log("Loading history...")
      const userStore = useUsersStore();
      const session_id = userStore.user.session_id;
      try {
        const response = await fetchWrapper.get(`${chatURL}/history/${session_id}`);
        if (response.length > 0) {
          this.messages = response
        }
        console.log("History:", this.messages);
      } catch (error) {
        console.error(error);
      } 
    },

    async sendPrompt(content) {      
      console.log("Sending prompt...")
      const userStore = useUsersStore();

      var msg = {
        sender: "human",
        content: content
      }

      this.messages = [...this.messages, msg];
      
      const data = {
        session_id: userStore.user.session_id,
        user_prompt: content
      }

      this.socket.send(JSON.stringify(data));
      // try {
      //   const response = await fetchWrapper.post(`${chatURL}/msg`, data);
      //   msg = {
      //     sender: "ai",
      //     content: response.content
      //   }
      //   this.messages = [...this.messages, msg];
      // } catch (error) {
      //   console.error(error);
      // }
    }
  }
})