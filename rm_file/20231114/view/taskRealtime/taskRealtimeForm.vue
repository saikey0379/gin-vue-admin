<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="Name:" prop="name">
          <el-input v-model="formData.name" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="负责人:" prop="manager">
          <el-input v-model="formData.manager" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="描述:" prop="describe">
          <el-input v-model="formData.describe" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="节点:" prop="nodes">
          <el-input v-model="formData.nodes" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文件类型:" prop="fileType">
          <el-select v-model="formData.fileType" placeholder="请选择" :clearable="false">
            <el-option v-for="(item,key) in FileTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="文件ID:" prop="fileId">
          <el-input v-model.number="formData.fileId" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文件权限:" prop="fileMod">
          <el-input v-model="formData.fileMod" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="执行参数:" prop="parameter">
          <el-input v-model="formData.parameter" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="目标路径:" prop="destPath">
          <el-input v-model="formData.destPath" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createTaskRealtime,
  updateTaskRealtime,
  findTaskRealtime
} from '@/api/taskRealtime'

defineOptions({
    name: 'TaskRealtimeForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const FileTypeOptions = ref([])
const formData = ref({
            name: '',
            manager: '',
            describe: '',
            nodes: '',
            fileType: undefined,
            fileId: 0,
            fileMod: '',
            parameter: '',
            destPath: '',
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               describe : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               nodes : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               fileType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               fileId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findTaskRealtime({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.retaskRealtime
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    FileTypeOptions.value = await getDictFunc('FileType')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createTaskRealtime(formData.value)
               break
             case 'update':
               res = await updateTaskRealtime(formData.value)
               break
             default:
               res = await createTaskRealtime(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
