<template>
  <a-layout id="components-layout-demo-custom-trigger">
    <a-layout-sider v-model:collapsed="collapsed" :trigger="null" collapsible>
      <div class="logo"/>
      <a-menu theme="dark" mode="inline" v-model:selectedKeys="selectedKeys">
        
        <a-sub-menu key="user">
            <template v-slot:title>
              <span><user-outlined/>用户管理</span>
            </template>
            
              <a-menu-item key="1">
                <router-link to="/main/user_list">
                  用户列表
                </router-link>
              </a-menu-item>
            
          </a-sub-menu>

          <a-sub-menu key="product">
            <template v-slot:title>
              <span><shopping-outlined />商品管理</span>
            </template>
            
            <a-menu-item key="5">
              <router-link to="/main/product_list">
                商品列表
              </router-link>
            </a-menu-item>

            <a-menu-item key="6">
              <router-link to="/main/seckill_list">
                活动管理
              </router-link>
            </a-menu-item>
          </a-sub-menu>
      </a-menu>
    </a-layout-sider>


    <a-layout style="padding: 0 24px 24px">
        <a-layout-header style="background: #fff; padding: 0">
            <menu-unfold-outlined
            v-if="collapsed"
            class="trigger"
            @click="() => (collapsed = !collapsed)"/>
            <menu-fold-outlined v-else class="trigger" @click="() => (collapsed = !collapsed)" />
        </a-layout-header>


        <a-layout-content :style="{ background: '#fff', padding: '24px', margin: 0, minHeight: '380px' }">
          <!-- 显示右侧内容 -->
            <!-- 欢迎使用秒杀商城 -->
            <router-view/>
        </a-layout-content>
        
      </a-layout>

  </a-layout>
  
</template>
<script>
import {
  UserOutlined,
  ShoppingOutlined,
//   VideoCameraOutlined,
//   UploadOutlined,
  MenuUnfoldOutlined,
  MenuFoldOutlined,

} from '@ant-design/icons-vue';

export default {
  components: {
    UserOutlined,
    ShoppingOutlined,
    // VideoCameraOutlined,
    // UploadOutlined,
    MenuUnfoldOutlined,
    MenuFoldOutlined,
  },
  data() {
    return {
      selectedKeys: ['1'],
      collapsed: false,
    };
  },
};
</script>
<style lang="scss" scoped>

#components-layout-demo-custom-trigger .trigger {
  font-size: 18px;
  line-height: 64px;
  padding: 0 24px;
  cursor: pointer;
  transition: color 0.3s;
}

#components-layout-demo-custom-trigger .trigger:hover {
  color: #1890ff;
}

#components-layout-demo-custom-trigger .logo {
  height: 32px;
  background: rgba(255, 255, 255, 0.2) url("../assets/logo.png");
  background-size:168px 32px;
	background-repeat:no-repeat;
  margin: 16px;
}
</style>
