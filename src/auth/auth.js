
const front_token_key = "front_token"
const front_user_name_key = "front_user_name"

const admin_token_key = "admin_token"
const admin_user_name_key = "admin_user_name"


class Auth {

    constructor(){
        // this.front_token = null
        // this.front_username = null
        this.front_token = localStorage.getItem(front_token_key)?localStorage.getItem(front_token_key):""
        this.front_username = localStorage.getItem(front_user_name_key)?localStorage.getItem(front_user_name_key):""
        this.admin_token = localStorage.getItem(admin_token_key)?localStorage.getItem(admin_token_key):""
        this.admin_user_name = localStorage.getItem(admin_user_name_key)?localStorage.getItem(admin_user_name_key):""
    }

    // 缓存token、用户信息
    setFrontAuth(token,username){
        this.front_token = token
        this.front_username = username
        localStorage.setItem(front_token_key,token)
        localStorage.setItem(front_user_name_key,username)

    }

    setAdminAuth(token,username){
        this.admin_token = token
        this.admin_user_name = username
        localStorage.setItem(admin_token_key,token)
        localStorage.setItem(admin_user_name_key,username)

    }

    // 删除token、用户信息，用来做用户登出
    delFrontAuth(){
        this.front_token = null
        this.front_username = null
        localStorage.removeItem(front_token_key)
        localStorage.removeItem(front_user_name_key)


    }

    delAdminAuth(){
        this.admin_token = null
        this.admin_user_name = null
        localStorage.removeItem(admin_token_key)
        localStorage.removeItem(admin_user_name_key)

    }


}

// 单例
const auth = new Auth()
export default auth