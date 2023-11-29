<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="序列号:" prop="sn">
          <el-input v-model="formData.sn" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="资产编号:" prop="asset_id">
          <el-input v-model="formData.asset_id" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="主机名:" prop="hostname">
          <el-input v-model="formData.hostname" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="IP地址:" prop="ip">
          <el-input v-model="formData.ip" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="管理IP:" prop="manage_ip">
          <el-input v-model="formData.manage_ip" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="网段ID:" prop="network_id">
          <el-input v-model.number="formData.network_id" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="管理网段ID:" prop="manage_network_id">
          <el-input v-model.number="formData.manage_network_id" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="硬件配置模版:" prop="hardware_id">
          <el-select v-model="formData.hardware_id" placeholder="请选择" :clearable="false">
            <el-option v-for="(item,key) in intOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="PXE模版:" prop="pxe_id">
          <el-select v-model="formData.pxe_id" placeholder="请选择" :clearable="false">
            <el-option v-for="(item,key) in intOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="Kickstart模版:" prop="kickstart_id">
          <el-select v-model="formData.kickstart_id" placeholder="请选择" :clearable="false">
            <el-option v-for="(item,key) in intOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="机柜ID:" prop="cabinet_id">
          <el-select v-model="formData.cabinet_id" placeholder="请选择" :clearable="false">
            <el-option v-for="(item,key) in intOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="机柜U位:" prop="cabinet_u">
          <el-input v-model="formData.cabinet_u" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="设备状态:" prop="status">
          <el-input v-model="formData.status" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="开发负责人:" prop="manager_dev">
          <el-input v-model="formData.manager_dev" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="运维负责人:" prop="manager_ops">
          <el-input v-model="formData.manager_ops" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="设备标签ID:" prop="label_id">
          <el-select v-model="formData.label_id" placeholder="请选择" :clearable="false">
            <el-option v-for="(item,key) in intOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="设备描述:" prop="device_describe">
          <el-input v-model="formData.device_describe" :clearable="false" placeholder="请输入" />
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
  createBareMetal,
  updateBareMetal,
  findBareMetal
} from '@/api/sysDeviceBareMetal'

defineOptions({
    name: 'BareMetalForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const intOptions = ref([])
const formData = ref({
            sn: '',
            asset_id: '',
            hostname: '',
            ip: '',
            manage_ip: '',
            network_id: 0,
            manage_network_id: 0,
            hardware_id: undefined,
            pxe_id: undefined,
            kickstart_id: undefined,
            cabinet_id: undefined,
            cabinet_u: '',
            status: '',
            manager_dev: '',
            manager_ops: '',
            label_id: undefined,
            device_describe: '',
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
               manage_ip : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               network_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               manage_network_id : [{
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
               label_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               device_describe : [{
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
      const res = await findBareMetal({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reDeviceBareMetal
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    intOptions.value = await getDictFunc('int')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createBareMetal(formData.value)
               break
             case 'update':
               res = await updateBareMetal(formData.value)
               break
             default:
               res = await createBareMetal(formData.value)
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
