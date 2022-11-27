<template>
  <div>
    <a-button @click="ToAdd">添加</a-button>
    <a-table :columns="columns" :data-source="seckills" :scroll="{ x: 1300 }" :pagination="user_pagenation" @change="pageChange" bordered>
      <template #action="{record}">
        <!-- <a @click="delProduct(record)">删除</a> -->

        <a-popconfirm v-if="seckills.length" title="确认删除吗?" @confirm="delSeckill(record)">
          <a-button>删除</a-button>
        </a-popconfirm>

        <router-link :to="{name:'SeckillrEdit',params:{id:record.id}}"><a-button>编辑</a-button></router-link>
      </template>

    </a-table>
  </div>
</template>
<script>

import qs from "qs"
const columns = [
  { title: '活动Id', width: 80, dataIndex: 'id', key: 'id', fixed: 'left' },
  { title: '活动名称', width: 100, dataIndex: 'name', key: 'name', fixed: 'left' },
  { title: '商品名称', width: 100, dataIndex: 'pname', key: 'pname', fixed: 'left' },
  { title: '活动价格', dataIndex: 'price', key: 'price' },
  { title: '活动数量', dataIndex: 'num', key: 'num' },
  { title: '活动开始时间', dataIndex: 'start_time', key: 'start_time' },
  { title: '活动结束时间', dataIndex: 'end_time', key: 'end_time' },
  { title: '创建时间', dataIndex: 'create_time', key: 'create_time' },
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
      seckills:[],
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
        this.getSeckills(this.user_pagenation.current,this.user_pagenation.pageSize)
    },

    methods:{

      delSeckill(record){
        this.$axios.post("/seckill/seckill_del",qs.stringify({
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
        this.$router.push("/main/seckill_add")
      },

      pageChange(pagination){

        this.user_pagenation.current = pagination.current
        this.user_pagenation.pageSize = pagination.pageSize
        this.getSeckills(pagination.current,pagination.pageSize)
        

      },
        
        getSeckills(currentPage,pageSize){
          
            this.$axios.get("/seckill/get_seckill_list",{
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
                this.seckills = rep.data.seckills
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
