<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="序列号:" prop="sn">
          <el-input v-model="formData.sn" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="主机名:" prop="hostname">
          <el-input v-model="formData.hostname" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="IP地址:" prop="ip">
          <el-input v-model="formData.ip" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="网段ID:" prop="network_id">
          <el-input v-model.number="formData.network_id" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="硬件配置:" prop="hardware_id">
          <el-input v-model.number="formData.hardware_id" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="PXE:" prop="pxe_id">
          <el-input v-model.number="formData.pxe_id" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="Kickstart:" prop="kickstart_id">
          <el-input v-model.number="formData.kickstart_id" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="设备状态:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择" :clearable="false">
            <el-option v-for="(item,key) in DeviceStatusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="任务负责人:" prop="manager">
          <el-input v-model="formData.manager" :clearable="false" placeholder="请输入" />
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
  createOsInstallQueue,
  updateOsInstallQueue,
  findOsInstallQueue
} from '@/api/osInstallQueue'

defineOptions({
    name: 'OsInstallQueueForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const DeviceStatusOptions = ref([])
const formData = ref({
            sn: '',
            hostname: '',
            ip: '',
            network_id: 0,
            hardware_id: 0,
            pxe_id: 0,
            kickstart_id: 0,
            status: undefined,
            manager: '',
        })
// 验证规则
const rule = reactive({
               sn : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               hostname : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               ip : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               network_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               hardware_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               pxe_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               kickstart_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               manager : [{
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
      const res = await findOsInstallQueue({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reosInstallQueue
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    DeviceStatusOptions.value = await getDictFunc('DeviceStatus')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createOsInstallQueue(formData.value)
               break
             case 'update':
               res = await updateOsInstallQueue(formData.value)
               break
             default:
               res = await createOsInstallQueue(formData.value)
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
