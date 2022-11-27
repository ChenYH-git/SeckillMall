<template>
  <div class="page-container">
    <div class="page-title">
      <span class="left">商品抢购系统</span>

      <span class="right" v-if="$auth.front_username == ''">
        <router-link to="/login">
          登录
        </router-link >/
        <router-link to="/register">
          注册
        </router-link >
      </span>

      <span class="right" v-else>
        <router-link to="/login">
          欢迎：{{$auth.front_username}}
        </router-link >/
        <!-- <span><input type="button" @click="Logout" value="退出"/></span> -->
        <span style="font-size:18px"><a @click="Logout">登出</a></span>
      </span>


    </div>

    
    <ul class="goods-list" >
      <li class="goods-box" v-for="seckill in seckills" :key="seckill">
        <router-link :to="{name:'GoodsDetail',params:{id:seckill.id}}" class="inner-box">
          <div class="thumbnail">
            <img :src="'data:image/jpeg;base64,'+ seckill.pic" alt="">
          </div>
          <div class="goods-info">
            <div class="title"><b>{{seckill.name}}</b></div>

            <div class="price">
              <b>秒杀价：￥{{seckill.price}}</b> &nbsp;
              <del>原价:￥{{seckill.p_price}}</del>
            </div>

            <div class="desc">{{seckill.pdesc}}</div>
          </div>
        </router-link>
      </li>

    </ul>

    <a @click="load_data">
      <div class="load_more">
        加载更多
      </div>
    </a>
    
        
  </div>


  
</template>

<script>
export default {
  name: 'GoodsList',

  data(){
    return {
      current_page:1,
      page_size:8,
      seckills:[],
      total_page:undefined
    }
  },

  mounted(){

    this.getSeckillList(this.current_page,this.page_size)
  },


  methods:{
    getSeckillList(current_page,page_size){
      this.$axios.get("/seckill/front/get_seckill_list",{
        params:{
          currentPage:current_page,
          pageSize:page_size
        }
      }).then((rep)=> {

        this.seckills = this.seckills.concat(rep.data.seckill_list) 
        this.current_page = rep.data.current
        this.page_size = rep.data.page_size
        this.total_page = rep.data.total_page


      }).catch((rep) => {
        console.log(rep);
      })


    },

    load_data(){

      if (this.total_page >= (this.current_page+1)){
        this.getSeckillList(this.current_page+1,this.page_size)
      }else{
        alert("无更多数据")

      }
      
      

    },

    
    Logout(){
      this.$auth.delFrontAuth()
      this.$router.push("/login")
    }
  }
}
</script>


<style scoped lang="scss">

.load_more {
  text-align: center;
  margin-top: 30px;
  background-color:rgb(194, 212, 235);
}

.page-title {
  margin-left: -100px;
  margin-bottom: 50px;
  .left {
    font-size: 32px;;
  }
  .right {
    float:right;
    font-size: 18px;;
  }
}
.goods-list{
  display: flex;
  justify-content: flex-start;
  flex-wrap: wrap;
  .goods-box{
    width: 25%;
    height: 357px;
    padding: 4px;
    cursor: pointer;
    box-sizing: border-box;
    border: 1px solid #fff;
    &:hover{
      background-color: #f4f0ea;
      border: 1px solid #dcdcdc;
    }
    .inner-box{
      .thumbnail{
        img{
          width: 100%;
          height: 240px;
        }
      }
      .goods-info{
        padding: 10px;
        .price{
          color: brown;
          padding: 10px 0;
          text-align: center;
          font-size: 16px;
          del {
            text-decoration: line-through;
            font-size: 14px;
          }
        }
        .desc{
          color: #666;
        }
      }
    }
  }
}
</style>
