import $ from 'jquery'

export default {
    state: {
        id: "",
        username: "",
        photo: "",
        token: "",
        is_login: false,
    },
    getters: {

    },
    mutations: {
        updateUser(state, user) {
            state.id = user.id;
            state.username = user.username;
            state.photo = user.photo;
            state.is_login = user.is_login;
        },
        updateToken(state, token) {
            state.token = token
        },
        logout(state) {
            state.id = "";
            state.username = "";
            state.photo = "";
            state.token = "";
            state.is_login = false;
        }
    },
    actions: {
        login(context, data) {
            $.ajax({
                url: "http://127.0.0.1:8081/user/login",
                type: "get",
                data: {
                    username: data.username,
                    password: data.password,
                },
                success(resp) {
                    if (resp.error_message === "success") {
                        context.commit("updateToken", resp.token);
                        data.success(resp);
                    } else {
                        data.error(resp);
                    }
                },
                error(resp) {
                    data.error(resp);
                }
            });
        },
        getInfo(context, data) {
            $.ajax({
                url: "http://127.0.0.1:8081/auth/getinfo",
                type: "get",
                headers: {
                    Authorization: context.state.token,
                },
                success(resp) {
                    if (resp.error_message === "success") {
                        console.log("114514" + resp.data.Username);
                        context.commit("updateUser", {
                            id: resp.data.ID,
                            username: resp.data.Username,
                            photo: resp.data.Photo,
                            token: resp.data.AccessToken,
                            is_login: true
                        });
                        data.success(resp);
                    } else {
                        data.error(resp);
                    }
                },
                error(resp) {
                    data.error(resp);
                }
            })
        },
        register(context, data) {
            $.ajax({
                url: "http://127.0.0.1:8081/user/register",
                type: "post",
                data: {
                    username: data.username,
                    password: data.password,
                    repassword: data.repassword,
                },
                success(resp) {
                    if (resp.error_message === "success") {
                        context.commit("updateToken", resp.token);
                        data.success(resp);
                    } else {
                        data.error(resp);
                    }
                },
                error(resp) {
                    console.log(resp);
                }
            })
        },
        logout(context) {
            context.commit("logout");
        }
    },
    modules: {

    }
}