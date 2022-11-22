<template>
 <div class="container">
		<header>
			<nav>
				<ul>
					<li><a href="#">Gmail</a></li>
					<li><a href="#">Imágenes</a></li>
					<li><a href="#"><i class="fa fa-th"></i></a></li>
					<li><a href="#"><i class="fa fa-user-circle"></i></a></li>
				</ul>
			</nav>
		</header>

		<section>
			<article class="logo">
				<img src="../assets/EnronLogo.png" alt="Remake Google">
			</article>
			<form class="search" @submit.prevent="search">
				<div class="search__box">
					<span><i class="fa fa-search"></i></span>
					<input type="text"  v-model="keyword">
					<span><i class="fa fa-microphone"></i></span>
				</div>
				<div class="btns">
					
				
						<input type="submit" value="Google Search" class="btn-google" />
			
				</div>
			</form>
     
		</section>

		<footer>
			<div class="country">
				<p>El Salvador</p>
			</div>
			<div class="links">
				<ul>
					<li><a href="#">Publicidad</a></li>
					<li><a href="#">Negocios</a></li>
					<li><a href="#">Sobre Google</a></li>
					<li><a href="#">Cómo funciona la Búsqueda</a></li>
				</ul>
				<ul>
					<li><a href="#">Privacidad</a></li>
					<li><a href="#">Condiciones</a></li>
					<li><a href="#">Preferencias</a></li>
				</ul>
			</div>
		</footer>
	
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
