<template>
 
	<div class="flex h-screen flex-col items-center justify-center bg-white">
  <!--  -->
  <div>
    <img class=" w-[273px]" src="../assets/EnronLogo.png" alt="google Logo" />
  </div>
  <!--  -->
  <form class="search" @submit.prevent="search">
  <div class="md:w-[584px] mx-auto mt-7 flex w-[92%] items-center rounded-full border hover:shadow-md">
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
  </div>
  <div class="mt-3 flex justify-center  space-x-12">
					
				
					<input type="submit" value="Google Search" class="bg-[#f8f9fa] px-2 py-1 cursor-pointer" />
		
			</div>
  <!--  -->
  
</form>
  <!--  -->
</div>
</template>

<script>
export default {
  name: 'HomePage',
  data() {
    return {
      keyword: "",
    };
  },
  methods: {
    search() {
      const newQuery = this.keyword.split(" ").join("-");
      this.$router.push({
        name: "PageDetails",
        params: { query: newQuery },
      });
    },
  },
  
  }
  
  
  

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

*{
	margin: 0;
	padding: 0;
	box-sizing: border-box;
}
body{
	font-family: Arial, Helvetica, sans-serif;
	
}
a{
	text-decoration: none;
}
ul{
	list-style: none;
}
.container{
	background: white;
}
header{
	width: 100%;
	height: 63px;
}
header nav{
	height: 100%;
	display: flex;
	align-items: center;
	justify-content: flex-end;
}
header nav ul{
	width: 250px;
	display: flex;
	align-items: center;
	justify-content: center;
}
header nav ul li a{
	margin-right: 15px;
	display: flex;
	font-size: 13px;
	color: #444;
}
header nav ul li:nth-child(3) a,
header nav ul li:nth-child(4) a{
	margin-left: 10px;
	margin-right: 15px;
	font-size: 18px;
	color: #777;
}
header nav ul li:nth-child(4) a{
	font-size: 23px;
}
section{

	text-align: center;
}
.logo{
	width: 592px;
	margin: 0 auto;
	margin-bottom: 27px;
}
.logo img{
	width: 272px;
}
.search{
	width: 592px;
	margin: 0 auto;
}
.search__box{
	width: 582px;
	border-radius: 100px;
	border: 1px solid #ddd;
	display: flex;
	align-items: center;
	justify-content: center;
	margin: 0 auto;
	margin-bottom: 27px;
}
.search__box input{
	width: 510px;
	height: 44px;
	border: none;
	outline: none;
	padding: 18px;
	color: #444;
	font-size: 16px;
}
.search__box:hover{
	box-shadow: -1px 1px 6px #bbb;
	border-color: #ddd;
}
.search__box .fa-search{
	font-size: 15px;
	color: #aaa;
}
.search__box .fa-microphone{
	color: #4285f4;
	font-size: 18px;
}

.btns{
	width: 100%;
	text-align: center;
	margin-bottom: 27px;
}
.btns input{
	padding: 10px 15px;
	margin: 0 2px;
	border: 1px solid #fff;
	border-radius: 5px;
	color: #5f6368;
	font-size: 14px;
	outline: none;
}
.btns input:hover{
	border-color: #bbb;
	box-shadow: 0 1px 1px #ddd;
}
footer a:hover{
	text-decoration: underline;
}
footer{
	width: 100%;
	height: 82px;
	position: absolute;
	bottom: 0;
	background: #f2f2f2;
}
footer .country,
footer .links{
	color: rgba(0,0,0,.54);
	font-size: 15px;
	height: 50%;
	padding: 0 30px;
	display: flex;
	align-items: center;
}
footer .country{
	justify-content: flex-start;
	border: 1px solid #e4e4e4;
}
footer .links{
	justify-content: space-between;
}
footer ul li{
	display: inline;
}
footer .links a{
		color: #5f6368;
		margin-right: 23px;
		font-size: 14px;
}
footer .links ul li:last-child a{
	margin: 0;
}
.modal{
        width: 400px;
        padding: 20px;
        margin: 100px auto;
        background: white;
        border-radius: 10px;
    }
    .backdrop {
        top:0;
        position: fixed;
        background: rgba(0,0,0,0.5);
        width: 100%;
        height: 100%;
    }
    .modal h1{
        color:rgb(92, 92, 93);
    }
    .modal p{
        font-style: normal;
    }
    .sale{
        background:crimson;
        color: #fff;
    }
    .sale h1{
        color: #fff;
    }


</style>
