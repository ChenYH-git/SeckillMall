<template>
  <a-table :columns="columns" :data-source="front_users" :scroll="{ x: 1300 }" :pagination="user_pagenation" @change="pageChange" bordered>
    <template #action="{}">
      <a>启用</a>
      
    </template>
  </a-table>
</template>
<script>
const columns = [
  { title: '邮箱', width: 100, dataIndex: 'email', key: 'name', fixed: 'left' },
  { title: '描述', width: 100, dataIndex: 'desc', key: 'age', fixed: 'left' },
  { title: '状态', dataIndex: 'status', key: '1' },
  { title: '创建时间', dataIndex: 'create_time', key: '2' },
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
      front_users:[],
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
        this.getFrontUsers(this.user_pagenation.current,this.user_pagenation.pageSize)
    },

    methods:{

      pageChange(pagination){

        this.user_pagenation.current = pagination.current
        this.user_pagenation.pageSize = pagination.pageSize
        this.getFrontUsers(pagination.current,pagination.pageSize)
        

      },
        
        getFrontUsers(currentPage,pageSize){
            this.$axios.get("/user/get_front_users",{
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
                this.front_users = rep.data.front_users
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
