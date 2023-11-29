<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="SN:" prop="sn">
          <el-input v-model="formData.sn" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="资产编号:" prop="asset_id">
          <el-input v-model="formData.asset_id" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="实例名:" prop="hostname">
          <el-input v-model="formData.hostname" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="业务IP:" prop="ip">
          <el-input v-model="formData.ip" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="业务网络:" prop="network_id">
          <el-input v-model.number="formData.network_id" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="管理IP:" prop="manage_ip">
          <el-input v-model="formData.manage_ip" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="管理网络:" prop="manage_network_id">
          <el-input v-model.number="formData.manage_network_id" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="硬件配置:" prop="hardware_id">
          <el-input v-model.number="formData.hardware_id" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="系统:" prop="kickstart_id">
          <el-input v-model.number="formData.kickstart_id" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="机柜:" prop="cabinet_id">
          <el-input v-model.number="formData.cabinet_id" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="U位:" prop="cabinet_u">
          <el-input v-model="formData.cabinet_u" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择" :clearable="false">
            <el-option v-for="(item,key) in DeviceStatusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="开发:" prop="manager_dev">
          <el-input v-model="formData.manager_dev" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="运维:" prop="manager_ops">
          <el-input v-model="formData.manager_ops" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="标签:" prop="label">
          <el-input v-model="formData.label" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="描述:" prop="describe">
          <el-input v-model="formData.describe" :clearable="false" placeholder="请输入" />
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
  createDeviceBareMetal,
  updateDeviceBareMetal,
  findDeviceBareMetal
} from '@/api/deviceBareMetal'

defineOptions({
    name: 'DeviceBareMetalForm'
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
            asset_id: '',
            hostname: '',
            ip: '',
            network_id: 0,
            manage_ip: '',
            manage_network_id: 0,
            hardware_id: 0,
            kickstart_id: 0,
            cabinet_id: 0,
            cabinet_u: '',
            status: undefined,
            manager_dev: '',
            manager_ops: '',
            label: '',
            describe: '',
        })
// 验证规则
const rule = reactive({
               sn : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               asset_id : [{
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
               manage_ip : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               manage_network_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               kickstart_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               cabinet_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               cabinet_u : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               manager_dev : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               manager_ops : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               label : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               describe : [{
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
      const res = await findDeviceBareMetal({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.redeviceBareMetal
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
               res = await createDeviceBareMetal(formData.value)
               break
             case 'update':
               res = await updateDeviceBareMetal(formData.value)
               break
             default:
               res = await createDeviceBareMetal(formData.value)
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
