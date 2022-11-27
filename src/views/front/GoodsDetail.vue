<template>
  <div class="page-container">
    <h1 class="page-title">商品详情页</h1>
    <div class="main-body">
      <div class="goods-detail">
        <div class="picture-box">
          <img :src="'data:image/jpeg;base64,'+ pic" alt="">
        </div>
        <div class="goods-info">
          <h1 class="title">{{name}}</h1>
          <div class="desc">{{pdesc}}</div>
          <ul class="text-info">
            <li>
              <span class="label">秒杀价</span>
              <span class="seckill-price">¥{{price}}</span>
              <span class="real-price">¥{{p_price}}</span>
            </li>
            <li>
              <span class="label">库存</span>
              <span class="value">{{num}}</span>
            </li>
            <li>
              <span class="label">单位</span>
              <span class="value">{{unit}}</span>
            </li>
            <li>
              <span class="label">秒杀时间</span>
              <span class="seckill-time">{{start_time}} -- {{end_time}}</span>
            </li>
            <li>
              <span class="label">抢购倒计时</span>
              <span class="seckill-time" style="color:red;font-size:18px;font-weight: 550;" v-if="message == ''">
                {{hours}}:{{minutes}}:{{seconds}}
              </span>
              <span class="seckill-time" style="color:red;font-size:18px;font-weight: 550;" v-else>
                {{message}}
              </span>
            </li>
          </ul>
          <div class="btn-box">
            <button class="submit-btn" type="primary" @click="onSubmit">立即购买</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import qs from "qs"
export default {
  name: "GoodsDetail",
  data(){
    return {
      id:undefined,
      name:"",
      num:undefined,
      price:undefined,
      pid:undefined,
      pic:"",
      p_price:"",
      pdesc:"",
      unit:"",
      start_time:"",
      end_time:"",

      hours:"",
      minutes:"",
      seconds:"",
      message:""


    }
  },
  mounted(){
    this.get_seckill_by_id();

    setInterval(() => {

      // 开始时间 - 当前时间  
      const start_time = Date.parse(this.start_time);
      const now_time = Date.parse(new Date());

      const sub = start_time - now_time;
      if (sub > 0) {
        const hours = Math.floor((sub % (24 * 60 * 60 * 1000)) / (60 * 60 * 1000)); 
        const minutes = Math.floor((sub % (60 * 60 * 1000)) / (60 * 1000)); 
        const seconds = Math.floor((sub % (60 * 1000)) / (1000)); 

        this.hours = hours
        this.minutes = minutes
        this.seconds = seconds
      }else{
        this.message = "抢购已开始"
      }
      
    }, 1000);
            
  },

  methods:{

    async getResult(){
      try{
        const rep = await this.$axios.get("/seckill/front/get_seckill_result",{
          params:{
          id:this.$route.params.id
        },
          headers:{
            Authorization:"Bearer " + localStorage.getItem("front_token")
          }
        })
        // console.log(rep);
        console.log(rep.data);
        console.log(rep.data.code);

        if(rep.data.code == 200 ){
          alert(rep.data.msg)
          window.location.reload()
        }else{
          setTimeout(() => {
            this.getResult()
          }, 2000);
        }
      }catch(err){
        alert(err)
      }

    },


    onSubmit(){

      const start_time = Date.parse(this.start_time);
      const now_time = Date.parse(new Date());

      const sub = start_time - now_time;

      if (sub > 0){
        alert("抢购还未开始")
        return
      }

    
      this.$axios.post("/seckill/front/seckill",qs.stringify({
        id:this.id
      }),{
        headers:{
          Authorization:"Bearer " + localStorage.getItem("front_token")
        }
      }).then((rep) => {
      
        if (rep.data.code == 200){
          alert(rep.data.msg)
          window.location.reload()
        }else {
          this.getResult()
          alert(rep.data.msg)


        }
      }).catch((rep) => {
        console.log(rep);

      })

    },
    get_seckill_by_id(){
      this.$axios.get("/seckill/front/seckill_detail",{
        params:{
          id:this.$route.params.id
        },
        headers:{
          Authorization:"Bearer " + localStorage.getItem("front_token")
        }
      }).then((rep) => {
        console.log(rep);
        if (rep.data.code == 200){
          this.id = rep.data.seckill.id
          this.name = rep.data.seckill.name
          this.num = rep.data.seckill.num
          this.price = rep.data.seckill.price
          this.pid = rep.data.seckill.pid
          this.pic = rep.data.seckill.pic
          this.p_price = rep.data.seckill.p_price
          this.pdesc = rep.data.seckill.pdesc
          this.unit = rep.data.seckill.unit
          this.start_time = rep.data.seckill.start_time
          this.end_time = rep.data.seckill.end_time

        }else{
          alert(rep.data.msg)
          console.log(rep.data.msg);
        }

      }).catch((rep) => {
        console.log(rep);
      })
    },

    
  }

}
</script>


<style scoped lang="scss">
.main-body{
  width: 100%;
  .goods-detail{
    display: flex;
    justify-content: flex-start;
    .picture-box{
      width: 380px;
      height: 380px;
      img{
        width: 100%;
        height: 100%;
      }
    }
    .goods-info{
      margin-left: 40px;
      flex: 1;
      text-align: left;
      .desc{
        padding: 5px 0;
        color: #666666;
        font-size: 13px;
      }
      .text-info{
        padding-top: 30px;
        li{
          display: flex;
          flex-direction: row;
          justify-content: flex-start;
          align-items: center;
          padding: 10px;
          .label{
            font-weight: 600;
            width: 100px;
          }
          .seckill-price{
            font-size: 26px;
            color: crimson;
            font-weight: 600;
          }
          .real-price{
            margin-left: 10px;
            font-size: 13px;
            text-decoration: line-through;
          }
        }
      }
      .btn-box{
        padding-top:40px;
        .submit-btn{
          color: #b4a078;
          border: 1px solid #b4a078;
          background-color: #f5f3ef;
          width: 168px;
          height: 50px;
          line-height: 50px;
          font-size: 18px;
          cursor: pointer;
        }
      }
    }
  }
}
</style>