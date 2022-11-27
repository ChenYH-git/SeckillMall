<template>
  <div>
    <a-button @click="ToAdd">添加</a-button>
    <a-table :columns="columns" :data-source="products" :scroll="{ x: 1300 }" :pagination="user_pagenation" @change="pageChange" bordered>
      <template #action="{record}">
        <!-- <a @click="delProduct(record)">删除</a> -->

        <a-popconfirm v-if="products.length" title="确认删除吗?" @confirm="delProduct(record)">
          <a-button>删除</a-button>
        </a-popconfirm>

        <router-link :to="{name:'ProductEdit',params:{id:record.id}}"><a-button>编辑</a-button></router-link>
      </template>

    </a-table>
  </div>
</template>
<script>

import qs from "qs"
const columns = [
  { title: '商品Id', width: 80, dataIndex: 'id', key: 'id', fixed: 'left' },
  { title: '商品名称', width: 100, dataIndex: 'name', key: 'name', fixed: 'left' },
  { title: '价格', width: 100, dataIndex: 'price', key: 'price', fixed: 'left' },
  { title: '商品数量', dataIndex: 'num', key: '1' },
  { title: '商品单位', dataIndex: 'unit', key: '2' },
  { title: '商品图片', dataIndex: 'pic', key: '3' },
  { title: '商品描述', dataIndex: 'desc', key: '4' },
  { title: '创建时间', dataIndex: 'create_time', key: '5' },
  {
    title: '操作',
    key: 'operation',
    fixed: 'right',
    width: 100,
    slots: { customRender: 'action' },
  },
];



export default {
  data() {
    return {
      products:[],
      columns,

      user_pagenation:{
        current:1,
        pageSize:5,
        pageSizeOptions:['5','10','20'],
        showSizeChanger:true,
        total:0

      }
      
    };
  },

   mounted(){
        this.getProducts(this.user_pagenation.current,this.user_pagenation.pageSize)
    },

    methods:{

  
      delProduct(record){
        console.log(record.id);
        this.$axios.post("/product/product_del",qs.stringify({
          id:record.id,
         
        }),{
           headers:{
            Authorization:"Bearer " + localStorage.getItem("admin_token")
          }
        }).then((rep) => {
          console.log(rep);
          if (rep.data.code == 200){
            alert(rep.data.msg)
            // this.$router.push("/main/product_list")
            window.location.reload()

          }else {
            alert(rep.data.msg)
          }

        }).catch((rep) => {
          console.log(rep);
          
        })

      },

      ToAdd(){
        this.$router.push("/main/product_add")
      },

      pageChange(pagination){

        this.user_pagenation.current = pagination.current
        this.user_pagenation.pageSize = pagination.pageSize
        this.getProducts(pagination.current,pagination.pageSize)
        

      },
        
        getProducts(currentPage,pageSize){
          
            this.$axios.get("/product/get_product_list",{
              params:{
                currentPage:currentPage,
                pageSize:pageSize
              },
              headers:{
                Authorization:"Bearer " + localStorage.getItem("admin_token")
              }
            }).then((rep) => {

                if (rep.data.code != 200){
                  this.$router.push("/admin")
                }
                this.products = rep.data.products
                this.user_pagenation.current = rep.data.current_page
                this.user_pagenation.pageSize = rep.data.page_size
                this.user_pagenation.total = rep.data.total
                console.log(rep);

            }).catch((rep) => {
                console.log(rep);

            })

        }
    }

};
</script>
