<template>
  <div>
   <div class="result">
    
    <form class="md:w-[584px] mx-auto mt-7 flex w-[92%] items-center rounded-full border hover:shadow-md" 
    @submit.prevent="searchMessages">
    
      <div class="pl-5">
      <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
        <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
      </svg>
    </div>
	
    <input type="text" class="w-full bg-transparent rounded-full py-[14px] pl-4 outline-none"  v-model="keyword" />
    <div class="pr-5">
      <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
        <path stroke-linecap="round" stroke-linejoin="round" d="M19 11a7 7 0 01-7 7m0 0a7 7 0 01-7-7m7 7v4m0 0H8m4 0h4m-4-8a3 3 0 01-3-3V5a3 3 0 116 0v6a3 3 0 01-3 3z" />
      </svg>
    </div>
    </form>
   
    
    
   </div> 
  
    
    <div className="mx-auto w-full min-h-screen px-3 sm:pl-[5%] md:pl-[14%] lg:pl-52 font-open-sans dark:bg-primary-dark  dark:text-white">
      <p className="text-gray-600 dark:text-gray-400 text-md mb-3 mt-3">
        About 365 results
      </p>

     
          <div
          v-for="message in messages" :key="message"
            className="max-w-xl py-4 px-3 text-xs mb-4 shadow ring-gray-200 dark:ring-[#303134] dark:ring-1 dark:shadow-3xl ring-2 sm:ring-0 sm:text-base sm:shadow-none rounded-lg"
          >
            <div className="group">
              <p className="text-sml line-clamp-1">
                {{message._source.Subject}}
              </p>
              <a >
                <h2 v-on:click="showMessage(message._source.Message)" className="truncate text-xl text-blue-800 dark:text-blue-400 font-medium group-hover:underline cursor-pointer">
                  {{message._source.From}}
                </h2>
              </a>
            </div>

            <p className="line-clamp-2">{{message._source.To}}</p>
          </div>
       

      
    </div>
   
    
  
   
   
    
    <div class="w-full h-full bg-gray-900 bg-opacity-80 top-0 fixed sticky-0"  v-if="isShowModal"
    :title="'ModalTitle'">
    <div
	class="relative top-20 mx-auto p-5 border w-[52%] shadow-lg rounded-md bg-white"
  v-if="isShowModal"
>
	<div class="mt-3 text-center">
		<h1>{{title}}</h1>
    <textarea id="textarea1" rows="25" class="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" v-model="contentMessage" ref="update" ></textarea>
		<div class="items-center px-4 py-3">
			<button
				id="ok-btn"
        @click="close" 
				class="px-4 py-2 bg-blue-500 text-white text-base font-medium rounded-md w-full shadow-sm hover:bg-black-600 focus:outline-none focus:ring-2 focus:ring-black-300"
			>
				Cerrar
			</button>
		</div>
	</div>
</div>
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
    
    },
 
    
    }
    
  
  

</script>

