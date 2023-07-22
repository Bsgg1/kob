<template>
    <ContentField>
        <div class="row justify-content-md-center">
            <div class="col-3">
                <form @submit.prevent="register">
                    <div class="mb-3">
                    <label for="username" class="form-label">用户名</label>
                    <input v-model="username" type="text" class="form-control" id="username" placeholder="请输入用户名">
                    </div>
                    <div class="mb-3">
                    <label for="password" class="form-label">密码</label>
                    <input v-model="password" type="text" class="form-control" id="password" placeholder="请输入密码">
                    </div>
                    <div class="mb-3">
                    <label for="repassword" class="form-label">确认密码</label>
                    <input v-model="repassword" type="text" class="form-control" id="repassword" placeholder="请确认密码">
                    </div>
                    <div class="error-message">{{ error_message }}</div>
                    <button type="submit" class="btn btn-primary">提交</button>
                </form>
            </div>
        </div>
    </ContentField>

</template>

<script>
import ContentField from '@/components/ContentField.vue'
import { useStore } from 'vuex'
import { ref } from 'vue'
import router from '../../../router/index'
export default{
    components:{
        ContentField
    },
    setup() {
        const store =useStore();
        let username = ref("");
        let password = ref("");
        let repassword = ref("");
        let error_message = ref("")
        const register = ()=>{
            error_message.value = ""
            store.dispatch("register",{
                username: username.value,
                password: password.value,
                repassword: repassword.value,
                success() {
                    store.dispatch("getInfo",{
                        success() {
                            router.push({name:"home"})
                            console.log(store.state.user)
                        }
                    })
                },
                error() {
                    error_message.value = "抱歉，请重新输入用户名和密码";
                }
            })
        }
        return {
            username,
            password,
            repassword,
            error_message,
            register
        }
    }
}
</script>

<style scoped>
.error-message{
    color: red;
}
</style>