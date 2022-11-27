<template>
  <a-form :model="form" :label-col="labelCol" :wrapper-col="wrapperCol">

    <a-form-item label="商品名称">
      <a-input v-model:value="form.name" placeholder="请输入商品名称" />
    </a-form-item>

    <a-form-item label="商品价格">
      <a-input v-model:value="form.price" placeholder="请输入商品价格，保留两位小数"/>
    </a-form-item>

    <a-form-item label="商品库存">
      <a-input v-model:value="form.num" placeholder="请输入商品库存"/>
    </a-form-item>

    <a-form-item label="商品单位">
      <a-input v-model:value="form.unit" placeholder="请输入商品单位"/>
    </a-form-item>

    <a-form-item label="商品原图">
      <img :src="'data:image/jpeg;base64,'+ form.pic" alt="" style="width:300px;height:200px">
    </a-form-item>

    <a-form-item label="商品新图">
      <a-input type="file"/>
    </a-form-item>


    <a-form-item label="商品描述">
      <a-input v-model:value="form.desc" placeholder="请输入商品描述"/>
    </a-form-item>

    
    <a-form-item :wrapper-col="{ span: 14, offset: 4 }">
      <a-button type="primary" @click="DoEdit">
        提交
      </a-button>
      <a-button style="margin-left: 10px;" @click="BackList">
        返回
      </a-button>
    </a-form-item>
  </a-form>
</template>
<script>
export default {
  data() {
    return {
      labelCol: { span: 4 },
      wrapperCol: { span: 14 },
      form: {
        id:undefined,
        name: '',
        price: undefined,
        num: undefined,
        unit: '',
        desc: '',
        pic:"",
      },
    };
  },
  created(){
    this.get_product_by_id()

  },
  methods: {

    get_product_by_id(){
      this.$axios.get("/product/to_product_edit",{
        params:{
          id:this.$route.params.id
        },
        
        headers:{
          Authorization:"Bearer " + localStorage.getItem("admin_token")
        }
      
      }).then((rep) => {

        console.log(rep);

        if (rep.data.code == 200){
          this.form.id = rep.data.product.id
          this.form.name = rep.data.product.name
          this.form.price = rep.data.product.price
          this.form.num = rep.data.product.num
          this.form.unit = rep.data.product.unit
          this.form.desc = rep.data.product.desc
          this.form.pic = rep.data.img_base64
        }else {
          alert("请传正确的商品id")
        }

        


      }).catch((rep) => {
        console.log(rep);
      })
    },

      BackList(){
          this.$router.push("/main/product_list")

      },
    DoEdit() {

        const form_data = new FormData();
        form_data.append("id",this.form.id)
        form_data.append("name",this.form.name)
        form_data.append("price",this.form.price)
        form_data.append("num",this.form.num)
        form_data.append("unit",this.form.unit)
        form_data.append("pic",document.querySelector('input[type=file]').files[0])
        form_data.append("desc",this.form.desc)

        this.$axios.post("/product/do_product_edit",form_data,{
          headers:{
                Authorization:"Bearer " + localStorage.getItem("admin_token")
              }
        }).then((rep) => {
            console.log(rep);
            if (rep.data.code == 200){
                alert(rep.data.msg)
                this.$router.push("/main/product_list")
            }else {
                alert(rep.data.msg)
            }

        }).catch((rep) => {
            console.log(rep);

        })
    },
  },
};
</script>
