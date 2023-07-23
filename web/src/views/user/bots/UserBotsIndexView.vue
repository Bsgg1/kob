<template>
    <div class="container">
        <div class="row">
            <div class="col-3">
                <div class="card" style="margin-top: 20px;">
                    <div class="card-body">
                        <img :src="$store.state.user.photo" alt="" style="width: 100%;">
                    </div>
                </div>
            </div>
            <div class="col-9">
                <div class="card" style="margin-top: 20px;">
                    <div class="card-header">
                        <span style="font-size: 120%;">我的bot</span>
        
                        <button type="button" class="btn btn-primary float-end" data-bs-toggle="modal" data-bs-target="#add-bot-btn">
                            创建bot
                        </button>
                        <!-- Button trigger modal -->
                    

                    <!-- Modal -->
                    <div class="modal fade" id="add-bot-btn" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
                    <div class="modal-dialog modal-xl">
                        <div class="modal-content">
                        <div class="modal-header">
                            <h1 class="modal-title fs-5">创建bot</h1>
                            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                        </div>
                        <div class="modal-body">
                            <div class="mb-3">
                                <label for="add-bot-title" class="form-label">名称</label>
                                <input v-model="addbot.title" type="text" class="form-control" id="add-bot-title" placeholder="请输入bot名称">
                            </div>
                            <div class="mb-3">
                                <label for="add-bot-description" class="form-label">简介</label>
                                <textarea v-model="addbot.description" class="form-control" id="add-bot-description" rows="3" placeholder="请输入bot信息"></textarea>
                            </div>
                            <div class="mb-3">
                                <label for="add-bot-code" class="form-label">代码</label>
                                <VAceEditor v-model:value="addbot.content" @init="editorInit" lang="c_cpp"
                                                theme="textmate" style="height: 300px" :options="{
                                                    enableBasicAutocompletion: true, //启用基本自动完成
                                                    enableSnippets: true, // 启用代码段
                                                    enableLiveAutocompletion: true, // 启用实时自动完成
                                                    fontSize: 18, //设置字号
                                                    tabSize: 4, // 标签大小
                                                    showPrintMargin: false, //去除编辑器里的竖线
                                                    highlightActiveLine: true,
                                                }" />

                            </div>
                        </div>
                        <div class="modal-footer">
                            
                            <div class="error-message">{{ addbot.error_message }}</div>
                            <button type="button" class="btn btn-primary" @click="add">创建</button>
                            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">取消</button>
                        </div>
                        </div>
                    </div>
                    </div>
                    </div>
                    <div class="card-body">
                        <table class="table table-hover">
                            <thead>
                                <tr>
                                    <th>名称</th>
                                    <th>创建时间</th>
                                    <th>操作</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr v-for="bot in bots" :key="bot.id">
                                    <td>{{ bot.Title }}</td>
                                    <td>{{ bot.CreatedAt }}</td>
                                    <td>
                                        <button type="button" class="btn btn-secondary" data-bs-toggle="modal" :data-bs-target="'#update-modal-id-' + bot.ID">修改</button>
                                        <button type="button" class="btn btn-danger" @click="remove(bot)">删除</button>
                                        <div class="modal fade" :id="'update-modal-id-' + bot.ID" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
                                        <div class="modal-dialog modal-xl">
                                            <div class="modal-content">
                                            <div class="modal-header">
                                                <h1 class="modal-title fs-5">修改bot</h1>
                                                <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                                            </div>
                                            <div class="modal-body">
                                                <div class="mb-3">
                                                    <label for="add-bot-title" class="form-label">名称</label>
                                                    <input v-model="bot.Title" type="text" class="form-control" id="add-bot-title" placeholder="请输入bot名称">
                                                </div>
                                                <div class="mb-3">
                                                    <label for="add-bot-description" class="form-label">简介</label>
                                                    <textarea v-model="bot.Description" class="form-control" id="add-bot-description" rows="3" placeholder="请输入bot信息"></textarea>
                                                </div>
                                                <div class="mb-3">
                                                    <label for="add-bot-code" class="form-label">代码</label>
                                                    <VAceEditor v-model:value="bot.Content" @init="editorInit" lang="c_cpp"
                                                theme="textmate" style="height: 300px" :options="{
                                                    enableBasicAutocompletion: true, //启用基本自动完成
                                                    enableSnippets: true, // 启用代码段
                                                    enableLiveAutocompletion: true, // 启用实时自动完成
                                                    fontSize: 18, //设置字号
                                                    tabSize: 4, // 标签大小
                                                    showPrintMargin: false, //去除编辑器里的竖线
                                                    highlightActiveLine: true,
                                                }" />
                                                </div>
                                            </div>
                                            <div class="modal-footer">
                                                
                                                <div class="error-message">{{ addbot.error_message }}</div>
                                                <button type="button" class="btn btn-primary" @click="update(bot)">保存修改</button>
                                                <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">取消</button>
                                            </div>
                                            </div>
                                        </div>
                                        </div>
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { Modal } from 'bootstrap/dist/js/bootstrap'
import { ref,reactive } from 'vue'
import $ from 'jquery'
import { useStore } from 'vuex'
import { VAceEditor } from 'vue3-ace-editor';
import ace from 'ace-builds';
import 'ace-builds/src-noconflict/mode-c_cpp';
import 'ace-builds/src-noconflict/mode-json';
import 'ace-builds/src-noconflict/theme-chrome';
import 'ace-builds/src-noconflict/ext-language_tools';
export default{
    components:{
        VAceEditor,
    },
    setup () {
        ace.config.set(
            "basePath",
            "https://cdn.jsdelivr.net/npm/ace-builds@" +
            require("ace-builds").version +
            "/src-noconflict/")
        const store = useStore();
        let bots = ref([])
        const addbot = reactive({
            title: "",
            description: "",
            content: "",
            error_message: "",
        });
        const refresh_list = () => {
            $.ajax({
                url: "http://127.0.0.1:8081/auth/bot/list",
                type: "get",
                headers: {
                    Authorization: store.state.user.token,
                },
                success(resp) {
                    bots.value = resp.bots;
                    
                    console.log(resp);
                },
                error(resp) {
                    console.log(resp);
                },
            })
        }
        refresh_list();

        const add = ()=> {
            addbot.error_message = "";
            $.ajax({
                url: "http://127.0.0.1:8081/auth/bot/add",
                type: "post",
                data: {
                    title: addbot.title,
                    description: addbot.description,
                    content: addbot.content,
                },
                headers: {
                    Authorization: store.state.user.token,
                },
                success(resp) {
                    if(resp.error_message === "success") {
                        addbot.title = "";
                        addbot.description = "";
                        addbot.content = "";
                        Modal.getInstance("#add-bot-btn").hide();
                        refresh_list()
                    } else {
                        addbot.error_message = resp.error_message;
                    }
                }
            })
        }
        const remove = (bot)=> {
                $.ajax({
                    url: "http://127.0.0.1:8081/auth/bot/remove",
                    type: "post",
                    data: {
                        id: bot.ID,
                    },
                    headers: {
                        Authorization: store.state.user.token,
                    },
                    success(resp) {
                        if(resp.error_message === "success") {
                            refresh_list();
                        } 
                    }
                })
            }
            const update = (bot)=> {
            addbot.error_message = "";
            $.ajax({
                url: "http://127.0.0.1:8081/auth/bot/update",
                type: "post",
                data: {
                    id: bot.ID,
                    title: bot.Title,
                    description: bot.Description,
                    content: bot.Content,
                },
                headers: {
                    Authorization: store.state.user.token,
                },
                success(resp) {
                    if(resp.error_message === "success") {
                        Modal.getInstance('#update-modal-id-' + bot.ID).hide();
                        refresh_list()
                    } else {
                        addbot.error_message = resp.error_message;
                    }
                }
            })
        }
        return {
            bots,
            addbot,
            refresh_list,
            add,
            remove,
            update,
        }
    }
}
</script>

<style scoped>
.error-message{
    color: red;
}
</style>