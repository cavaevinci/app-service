import{L as s,N as t}from"./index-f3458156.js";const c="https://dev.asai.astrosynapse.ai/api",a=`${c}/tools`,l=s({id:"tool",state:()=>({records:{},record:{}}),actions:{async getTools(){try{const r=await t.get(`${a}`);this.records=r}catch(r){console.error(r)}},async getTool(r){try{const o=await t.get(`${a}/${r}`);this.record=o}catch(o){console.error(o)}},async saveAvatarTool(r){try{const o=await t.post(`${a}/save/avatar`,r);this.record=o}catch(o){console.error(o)}},async saveAgentTool(r){try{const o=await t.post(`${a}/save/agent`,r);this.record=o}catch(o){console.error(o)}},async toggleAvatarTool(r,o){try{await t.post(`${a}/${r}/toggle/avatar`,o)}catch(e){console.error(e)}},async toggleAgentTool(r,o){try{await t.post(`${a}/${r}/toggle/agent`,o)}catch(e){console.error(e)}}}});export{l as u};
