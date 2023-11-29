<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
        <el-form-item label="IP地址" prop="network">
         <el-input v-model="searchInfo.network" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="区域" prop="regionId">
            
             <el-input v-model.number="searchInfo.regionId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="网段ID" prop="segmentId">
            
             <el-input v-model.number="searchInfo.segmentId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="名称" prop="name">
         <el-input v-model="searchInfo.name" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="标签" prop="label">
         <el-input v-model="searchInfo.label" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" :disabled="!multipleSelection.length" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column sortable align="left" label="IP地址" prop="network" width="120" />
        <el-table-column sortable align="left" label="掩码" prop="netmask" width="120" />
        <el-table-column align="left" label="网关" prop="gateway" width="120" />
        <el-table-column sortable align="left" label="区域" prop="regionName" width="120" />
        <el-table-column sortable align="left" label="网段ID" prop="segmentName" width="120" />
        <el-table-column align="left" label="名称" prop="name" width="120" />
        <el-table-column sortable align="left" label="标签" prop="label" width="120" />
        <el-table-column align="left" label="描述" prop="describe" width="120" />
        <el-table-column align="left" label="操作">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateCmdbIpSubnetFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type==='create'?'添加':'修改'" destroy-on-close>
      <el-scrollbar height="500px">
          <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="IP地址:"  prop="network" >
              <el-input v-model="formData.network" :clearable="false"  placeholder="请输入IP地址" />
            </el-form-item>
            <el-form-item label="掩码:"  prop="netmask" >
              <el-input v-model.number="formData.netmask" :clearable="false" placeholder="请输入掩码" />
            </el-form-item>
            <el-form-item label="网关:"  prop="gateway" >
              <el-input v-model="formData.gateway" :clearable="false"  placeholder="请输入网关" />
            </el-form-item>
            <el-form-item label="区域:"  prop="regionId" >
              <el-select
                  v-model="formData.regionId"
                  style="width:256px"
                  placeholder="请选择区域"
                  clearable
                  @change="handleRegionSelect"
              >
                <el-option
                    v-for="item in regions"
                    :key="item.ID"
                    :value="item.ID"
                    :label="item.name"
                />
              </el-select>
              <el-icon
                  class="cursor-pointer ml-2 text-gray-600"
                  @click="getRegions"
              ><refresh /></el-icon>
              <el-icon
                  class="cursor-pointer ml-2 text-gray-600"
                  @click="goRegion"
              ><document-add /></el-icon>
            </el-form-item>
            <el-form-item label="网段ID:"  prop="segmentId" >
              <el-select
                  v-model="formData.segmentId"
                  style="width:256px"
                  :disabled="isOptionIpSegmentDisabled"
                  placeholder="请选择网段"
                  clearable
              >
                <el-option
                    v-for="item in ipSegments"
                    :key="item.ID"
                    :value="item.ID"
                    :label="item.name"
                />
              </el-select>
              <el-icon
                  class="cursor-pointer ml-2 text-gray-600"
                  @click="getCmdbIpSegmentsByRegionId"
              ><refresh /></el-icon>
              <el-icon
                  class="cursor-pointer ml-2 text-gray-600"
                  @click="goIpSegment"
              ><document-add /></el-icon>
            </el-form-item>
            <el-form-item v-if="showVLAN" label="VLAN:"  prop="vlanId" >
              <el-input v-model.number="formData.vlanId" :clearable="false" placeholder="请输入VLAN" />
            </el-form-item>
            <el-form-item label="名称:"  prop="name" >
              <el-input v-model="formData.name" :clearable="false"  placeholder="请输入名称" />
            </el-form-item>
            <el-form-item label="标签:"  prop="label" >
              <el-input v-model="formData.label" :clearable="false"  placeholder="请输入标签" />
            </el-form-item>
            <el-form-item label="描述:"  prop="describe" >
              <el-input v-model="formData.describe" :clearable="false"  placeholder="请输入描述" />
            </el-form-item>
          </el-form>
      </el-scrollbar>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="detailShow" style="width: 800px" lock-scroll :before-close="closeDetailShow" title="查看详情" destroy-on-close>
      <el-scrollbar height="550px">
        <el-descriptions column="1" border>
                <el-descriptions-item label="IP地址">
                        {{ formData.network }}
                </el-descriptions-item>
                <el-descriptions-item label="掩码">
                        {{ formData.netmask }}
                </el-descriptions-item>
                <el-descriptions-item label="网关">
                        {{ formatInfo("","",formData.gateway) }}
                </el-descriptions-item>
                <el-descriptions-item label="区域">
                        {{ formData.regionName }}
                </el-descriptions-item>
                <el-descriptions-item label="网段">
                        {{ formData.segmentName }}
                </el-descriptions-item>
                <el-descriptions-item label="VLAN">
                        {{ formatInfo("","",formData.vlanId) }}
                </el-descriptions-item>
                <el-descriptions-item label="名称">
                        {{ formData.name }}
                </el-descriptions-item>
                <el-descriptions-item label="标签">
                        {{ formData.label }}
                </el-descriptions-item>
                <el-descriptions-item label="描述">
                        {{ formData.describe }}
                </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createCmdbIpSubnet,
  deleteCmdbIpSubnet,
  deleteCmdbIpSubnetByIds,
  updateCmdbIpSubnet,
  findCmdbIpSubnet,
  getCmdbIpSubnetList
} from '@/api/cmdbIpSubnet'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, watch } from 'vue'
import { useRouter } from 'vue-router'
import { getCmdbRegionList } from '@/api/cmdbRegion'
import { getCmdbIpSegmentList} from '@/api/cmdbIpSegment'

defineOptions({
    name: 'CmdbIpSubnet'
})

const router = useRouter()

const goRegion = () => {
  router.push({ name: 'region' })
}

const goIpSegment = () => {
  router.push({ name: 'ipSegment' })
}

// 获取Region列表
const regions = ref([])
const getRegions = async() => {
  const res = await getCmdbRegionList()
  if (res.code === 0) {
    regions.value = res.data.list
  }
}

// Region处理
const handleRegionSelect = () => {
  getCmdbIpSegmentsByRegionId()
  resetShowVLAN()
};

// 根据RegionId获取网段列表
const ipSegments = ref([])
const isOptionIpSegmentDisabled = ref(true);

const getCmdbIpSegmentsByRegionId = () => {
  // 调用网段列表查询接口，并传递RegionId作为参数
  const { regionId } = formData.value;
  getCmdbIpSegmentList({ RegionId: regionId}).then(response => {
    // 根据接口返回的数据更新网段列表
    ipSegments.value = response.data.list;
    isOptionIpSegmentDisabled.value = regionId === 0;
  }).catch(error => {
    console.error("调用网段列表查询接口失败:", error);
  });
};

const formatInfo = (row, column, cellValue) => {
  if (cellValue === 0 ||cellValue === null || cellValue === "") {
    return '—';
  } else {
    return cellValue;
  }
};

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        network: '',
        gateway: '',
        name: '',
        label: '',
        describe: '',
        })

// 监听Region变化，调用网段列表查询接口
watch(formData, () => {
  getCmdbIpSegmentsByRegionId();
});

// 验证规则
const rule = reactive({
               network : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               netmask : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               regionId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               segmentId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               label : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               describe : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 排序
const sortChange = ({ prop, order }) => {
  searchInfo.value.sort = prop
  searchInfo.value.order = order
  getTableData()
}

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getCmdbIpSubnetList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteCmdbIpSubnetFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteCmdbIpSubnetByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateCmdbIpSubnetFunc = async(row) => {
    const res = await findCmdbIpSubnet({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.recmdbIpSubnet
        dialogFormVisible.value = true
        resetShowVLAN()
    }
}


// 删除行
const deleteCmdbIpSubnetFunc = async (row) => {
    const res = await deleteCmdbIpSubnet({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)


// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findCmdbIpSubnet({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.recmdbIpSubnet
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          network: '',
          gateway: '',
          name: '',
          label: '',
          describe: '',
          }
}


// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        network: '',
        gateway: '',
        name: '',
        label: '',
        describe: '',
        }
        showVLAN.value = false;
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createCmdbIpSubnet(formData.value)
                  break
                case 'update':
                  res = await updateCmdbIpSubnet(formData.value)
                  break
                default:
                  res = await createCmdbIpSubnet(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}

// 监听RegionID变化，确认是否展示VLAN输入框
const showVLAN = ref(false);

const resetShowVLAN = async() => {
  const { regionId } = formData.value;

  if (regions.value) {
    const region = regions.value.find(region => region.ID === regionId);

    if (region && region.regionType === 2) {
      showVLAN.value = true;
    } else {
      showVLAN.value = false;
      formData.value.vlanId = 0;
    }
  }
}

const init = () => {
  getRegions()
}

init()

</script>

<style>

</style>
