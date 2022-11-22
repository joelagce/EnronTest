<template>
  
   <div class="result">
    
    <form class="formContainer__form resultForm" @submit.prevent="searchMessages">
      <input type="text" class="formContainer__input" v-model="keyword" />
      <svg
        width="1.1em"
        height="1.1em"
        viewBox="0 0 16 16"
        class="bi bi-search searchIcon"
        fill="currentColor"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          fill-rule="evenodd"
          d="M10.442 10.442a1 1 0 0 1 1.415 0l3.85 3.85a1 1 0 0 1-1.414 1.415l-3.85-3.85a1 1 0 0 1 0-1.415z"
        />
        <path
          fill-rule="evenodd"
          d="M6.5 12a5.5 5.5 0 1 0 0-11 5.5 5.5 0 0 0 0 11zM13 6.5a6.5 6.5 0 1 1-13 0 6.5 6.5 0 0 1 13 0z"
        />
      </svg>
      <svg
        width="1.3em"
        height="1.3em"
        viewBox="0 0 16 16"
        class="bi bi-mic microphone"
        fill="currentColor"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          fill-rule="evenodd"
          d="M3.5 6.5A.5.5 0 0 1 4 7v1a4 4 0 0 0 8 0V7a.5.5 0 0 1 1 0v1a5 5 0 0 1-4.5 4.975V15h3a.5.5 0 0 1 0 1h-7a.5.5 0 0 1 0-1h3v-2.025A5 5 0 0 1 3 8V7a.5.5 0 0 1 .5-.5z"
        />
        <path
          fill-rule="evenodd"
          d="M10 8V3a2 2 0 1 0-4 0v5a2 2 0 1 0 4 0zM8 0a3 3 0 0 0-3 3v5a3 3 0 0 0 6 0V3a3 3 0 0 0-3-3z"
        />
      </svg>
    </form>
   
    
    
   </div> 
   <div class="area">
   <div class="show">
    <div class="show__results" v-for="message in messages" :key="message" >
      <p class="topLink" v-on:click="showMessage(message._source.Message)">{{message._source.Subject}}</p>
      <p class="link">
        <a >{{message._source.From}}</a>
      </p>
      <p class="description">
        {{message._source.To}}
      </p>
    </div>
    
  </div>
  
 
     
    </div>
   
    
    <div class="modal-backdrop"  v-if="isShowModal"
    :title="'ModalTitle'">
      <div class="modal">
        <header class="modal-header">
          <span class="modal__title">{{title}}</span>
          <img class="btn-close" @click="close"  src="@/assets/close.svg" alt="">
        </header>
        <section class="modal-body">
          <div class="input-field col s5">
        <textarea id="textarea1" class="materialize-textarea" v-model="contentMessage" ref="update" ></textarea>
      </div>
        </section>
       
      </div>
    </div>
    
</template>

<script>




export default {
 
  name: 'PageDetails',
  props: {
    msg: String,
   
  },
  
  
  data() {
    return {
      
      isShowModal: false,
      messages: [],
      googleResults: [],
      contentMessage: '',
      keyword: this.$route.params.query
        ? this.$route.params.query.split("-").join(" ")
        : "",
     
    
    }
  },
  methods: {
    close() {
      this.isShowModal = false
    },
    

   
    async searchMessages() {
      const newQuery = this.keyword.split(" ").join("-");
      this.$router.push({
        name: "PageDetails",
        params: { query: newQuery },
      });
    
      
    },
    showMessage(data){
      this.isShowModal = true
      this.contentMessage = data
      this.$refs.update.blur()
      console.log(this.$refs.update)
    },
	
	
  },
  async created() {
    if (this.$route.params.query) {
      try {
        const url = "http://localhost:3000/search"
      let options = {
        method: 'POST',
        headers: {
          'Access-Control-Allow-Origin' : '*',
          'Access-Control-Request-Method': 'POST, OPTIONS'
        }, 
        body: JSON.stringify({ input: this.keyword })
      }
      console.log(options)
      const resp = fetch(url, options)
      resp.then(res => res.json()).then(data => 
      data.hits.hits
      ).then( hits => hits.map((data) => this.messages.push(data)))
      } catch (error) {
        console.log(error);
      }
    }
    
    }
    }
  
  

</script>
<style>
.formContainer__input {
  width: 100%;
  border-radius: 24px;
  height: 50px;
  padding: 0 50px;
  border: 1px solid #dfe1ef;
  outline: none;
  font-size: 18px;
}
.formContainer__form {
  position: relative;
}
.searchIcon {
  position: absolute;
  left: 25px;
  top: 18px;
  color: #9aa0ab;
}
.microphone {
  position: absolute;
  right: 25px;
  top: 16px;
  color: #4285f4;
}
.result {
  width: 1200px;
  margin: 35px auto;
  display: flex;
  align-items: center;
}
p{
  cursor: pointer;
}
.result__img img {
  width: 92;
  height: 30px;
  cursor: pointer;
}
.resultForm {
  width: 600px;
  margin-left: 20px;
}
.resultForm .formContainer__input {
  width: 100% !important;
}
.show {
  width: 1200px;
  margin: 20px auto;
  
}
.show__results {
  width: 700px;
 
  margin-top: 20px;
}
.area{
  display: flex;
}
.topLink {
  color: #202124;
  font-size: 14px;
  margin-bottom: 6px;
}
.link {
  font-size: 20px;
  text-decoration: none;
  color: #609;
  margin-bottom: 10px;
}
.description {
  color: #4d5156;
  font-size: 14px;
}


</style>
<style lang="scss">
    .modal-backdrop {
      position: fixed;
      z-index: 999;
      top: 0;
      bottom: 0;
      left: 0;
      right: 0;
      background-color: rgba(0, 0, 0, 0.3);
      display: flex;
      justify-content: center;
      align-items: center;
    }
    .modal {
      position: relative;
      padding: 30px 30px 40px 30px;
      width: 630px;
      z-index: 1000;
      background: #FFFFFF;
      display: flex;
      flex-direction: column;
      border-radius: 8px;
      &__title {
        font-weight: 600;
        font-size: 24px;
        display: block;
        color: var(--black-title);
        margin-bottom: 30px;
      }
      &-footer {
        margin-top: 30px;
        display: flex;
        align-items: center;
      }
    }
    .btn-close {
      position: absolute;
      height: 20px;
      top: 17px;
      right: 17px;
      border: none;
      cursor: pointer;
      font-weight: bold;
      color: #2196F3;
      background: transparent;
    }
  </style>