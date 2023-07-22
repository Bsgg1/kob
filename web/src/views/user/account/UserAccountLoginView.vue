<template>
    <ContentField v-if="!$store.state.user.loading">
        <div class="row justify-content-md-center">
            <div class="col-3">
                <form @submit.prevent="login">
                    <div class="mb-3">
                    <label for="username" class="form-label">用户名</label>
                    <input v-model="name" type="text" class="form-control" id="username" placeholder="请输入用户名">
                    </div>
                    <div class="mb-3">
                    <label for="password" class="form-label">密码</label>
                    <input v-model="password" type="text" class="form-control" id="password" placeholder="请输入密码">
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
        
        const store = useStore();
        let name = ref('')
        let password = ref('')
        let error_message = ref('')
        
        const token = localStorage.getItem("jwt_token");
        if (token){
            store.commit("updateToken",token);
            store.dispatch("getInfo",{
                success() {
                    router.push({name:"home"});
                    store.commit("updateLoading",false);
                },
                error() {
                    store.commit("updateLoading",false);
                },
            })
        } else {
            store.commit("updateLoading",false);
        }
        const login =() => {
            error_message.value = "";
            store.dispatch("login",{
                username: name.value,
                password: password.value,
                success() {
                    store.dispatch("getInfo",{
                        success() {
                            router.push({name:"home"})
                            console.log(store.state.user)
                        }
                    })
                },
                error() {
                    error_message.value = "用户名或密码错误";
                }
            })
        }
        return {
            name,
            password,
            error_message,
            login,
        }
    }

}
</script>

<style scoped>
button {
    width: 100%;
}
.error-message{
    color: red;
}
</style>