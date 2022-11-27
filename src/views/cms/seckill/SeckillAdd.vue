<template>
  <a-form :model="form" :label-col="labelCol" :wrapper-col="wrapperCol">

    <a-form-item label="活动名称">
      <a-input v-model:value="form.name" placeholder="请输入活动名称" />
    </a-form-item>

    <a-form-item label="活动价格">
      <a-input v-model:value="form.price" placeholder="请输入活动价格，保留两位小，不能大于商品价格"/>
    </a-form-item>

    <a-form-item label="活动库存">
      <a-input v-model:value="form.num" placeholder="请输入活动库存"/>
    </a-form-item>

    <a-form-item label="选择关联商品">
      <a-select
        v-model:value="form.pid"
        show-search
        placeholder="请选择关联商品"
        option-filter-prop="children"
        :filter-option="filterOption"
        @focus="handleFocus"
        @blur="handleBlur"
        @change="handleChange"
        style="width: 300px"

        
    >
        <a-select-option v-for="product in form.products" :key="product" :value="product.id">
            {{product.pname}}
        </a-select-option>

    </a-select>
    </a-form-item>

    <a-form-item label="活动时间区间">
        <a-range-picker
    :disabled-date="disabledDate"
    :disabled-time="disabledRangeTime"
    :show-time="{
      hideDisabledOptions: true,
      defaultValue: [moment('00:00:00', 'HH:mm:ss'), moment('11:59:59', 'HH:mm:ss')],
    }"
    format="YYYY-MM-DD HH:mm:ss"

    valueFormat="YYYY-MM-DD HH:mm:ss"

    @ok="isok"
  />
    </a-form-item>

    
    <a-form-item :wrapper-col="{ span: 14, offset: 4 }">
      <a-button type="primary" @click="onSubmit">
        提交
      </a-button>
      <a-button style="margin-left: 10px;" @click="BackList">
        返回
      </a-button>
    </a-form-item>
  </a-form>
</template>
<script>

import moment from 'moment';
import qs from "qs"
export default {
  data() {
    return {
      labelCol: { span: 4 },
      wrapperCol: { span: 14 },
      form: {
        name: '',
        price: undefined,
        num: undefined,
        start_time: '',
        end_time: '',
        pid:undefined,
        products:[],

      },
    };
  },
  mounted(){
      this.getProduct()
  },
  methods: {



        moment,
    range(start, end) {
        const result = [];
        for (let i = start; i < end; i++) {
            result.push(i);
        }
        return result;
    },

    disabledDate(current) {
        // Can not select days before today and today
        return current && current < moment().endOf('day');
    },

    disabledDateTime() {
        return {
            disabledHours: () => this.range(0, 24).splice(4, 20),
            disabledMinutes: () => this.range(30, 60),
            disabledSeconds: () => [55, 56],
        };
    },

    disabledRangeTime(_, type) {
      if (type === 'start') {
        return {
          disabledHours: () => this.range(0, 60).splice(4, 20),
          disabledMinutes: () => this.range(30, 60),
          disabledSeconds: () => [55, 56],
        };
      }
      return {
        disabledHours: () => this.range(0, 60).splice(20, 4),
        disabledMinutes: () => this.range(0, 31),
        disabledSeconds: () => [55, 56],
      };
    },



      getProduct(){
          this.$axios.get("/seckill/get_products",{
              headers:{
                Authorization:"Bearer " + localStorage.getItem("admin_token")
              }
          }).then((rep) => {
              console.log(rep);
              this.form.products = rep.data.products

          }).catch((rep) => {
              console.log(rep);
          })

      },

      BackList(){
          this.$router.push("/main/seckill_list")

      },

    
    isok(values){
        console.log(values);
        this.form.start_time = values[0]
        this.form.end_time = values[1]

    },
    onSubmit() {
        console.log(this.form.pid.id);

        this.$axios.post("/seckill/seckill_add",qs.stringify({
            name:this.form.name,
            price:this.form.price,
            num:this.form.num,
            pid:this.form.pid,
            start_time:this.form.start_time,
            end_time:this.form.end_time,
            
        }),{
          headers:{
              Authorization:"Bearer " + localStorage.getItem("admin_token")
            }
        }).then((rep) => {
            console.log(rep);
            if (rep.data.code == 200){
                alert(rep.data.msg)
                this.$router.push("/main/seckill_list")

            }else{
                alert(rep.data.msg)
            }

        }).catch((rep) => {
            console.log(rep);
        })

        
    },
  },
};
</script>
