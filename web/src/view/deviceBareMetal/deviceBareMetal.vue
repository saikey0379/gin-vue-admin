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
        <el-form-item label="SN" prop="sn">
         <el-input v-model="searchInfo.sn" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="资产编号" prop="asset_id">
         <el-input v-model="searchInfo.asset_id" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="实例名" prop="hostname">
         <el-input v-model="searchInfo.hostname" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="业务IP" prop="ip">
         <el-input v-model="searchInfo.ip" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="业务网络" prop="network_id">
            
             <el-input v-model.number="searchInfo.network_id" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="管理IP" prop="manage_ip">
         <el-input v-model="searchInfo.manage_ip" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="管理网络" prop="manage_network_id">
            
             <el-input v-model.number="searchInfo.manage_network_id" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="硬件配置" prop="hardware_id">
            
             <el-input v-model.number="searchInfo.hardware_id" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="系统" prop="kickstart_id">
            
             <el-input v-model.number="searchInfo.kickstart_id" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="机柜" prop="cabinet_id">
            
             <el-input v-model.number="searchInfo.cabinet_id" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="开发" prop="manager_dev">
         <el-input v-model="searchInfo.manager_dev" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="运维" prop="manager_ops">
         <el-input v-model="searchInfo.manager_ops" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="标签" prop="label">
         <el-input v-model="searchInfo.label" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="描述" prop="describe">
         <el-input v-model="searchInfo.describe" placeholder="搜索条件" />

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
        <el-table-column align="left" label="SN" prop="sn" width="140" />
        <el-table-column sortable align="left" label="资产编号" prop="asset_id" width="120" />
        <el-table-column sortable align="left" label="实例名" prop="hostname" width="120" />
        <el-table-column sortable align="left" label="业务IP" prop="ip" width="120" />
        <el-table-column sortable align="left" label="状态" prop="status" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.status,DeviceStatusOptions) }}
            </template>
        </el-table-column>
        <el-table-column sortable align="left" label="开发" prop="manager_dev" width="120" />
        <el-table-column sortable align="left" label="运维" prop="manager_ops" width="120" />
        <el-table-column sortable align="left" label="标签" prop="label" width="120" />
        <el-table-column align="left" label="描述" prop="describe" width="120" />
        <el-table-column align="left" label="操作">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateDeviceBareMetalFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="plus"  @click="osInstall(scope.row)">重装系统</el-button>
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
            <el-form-item label="IDC:"  prop="idc_id" >
              <el-select
                  v-model="formData.idc_id"
                  style="width:256px"
                  placeholder="请选择IDC"
                  clearable
                  @change="handleIdcChange"
              >
                <el-option
                    v-for="item in idcs"
                    :key="item.ID"
                    :value="item.ID"
                    :label="item.name"
                />
              </el-select>
              <el-icon
                  class="cursor-pointer ml-2 text-gray-600"
                  @click="getIdcs"
              ><refresh /></el-icon>
              <el-icon
                  class="cursor-pointer ml-2 text-gray-600"
                  @click="goIdc"
              ><document-add /></el-icon>
            </el-form-item>
            <el-form-item label="SN:"  prop="sn" >
              <el-select
                  v-model="formData.sn"
                  style="width:256px"
                  placeholder="请选择纳管节点"
                  clearable
              >
                <el-option
                    v-for="item in sns"
                    :key="item.Sn"
                    :value="item.Sn"
                    :label="item.Sn"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="资产编号:"  prop="asset_id" >
              <el-input v-model="formData.asset_id" :clearable="false"  placeholder="请输入资产编号" />
            </el-form-item>
            <el-form-item label="实例名:"  prop="hostname" >
              <el-input v-model="formData.hostname" :clearable="false"  placeholder="请输入实例名" />
            </el-form-item>
            <el-form-item label="业务网段:"  prop="segmentIdService" >
              <el-select
                  v-model="formData.segmentIdService"
                  style="width:256px"
                  :disabled="isOptionIpSegmentServiceDisabled"
                  placeholder="请选择业务网段"
                  clearable
                  @change="getIdcIpSubnetServicesByIdcId"
              >
                <el-option
                    v-for="item in ipSegmentServices"
                    :key="item.ID"
                    :value="item.ID"
                    :label="item.name"
                />
              </el-select>
              <el-icon
                  class="cursor-pointer ml-2 text-gray-600"
                  @click="getIdcIpSegmentServicesByIdcId"
              ><refresh /></el-icon>
              <el-icon
                  class="cursor-pointer ml-2 text-gray-600"
                  @click="goIpSegment"
              ><document-add /></el-icon>
            </el-form-item>
            <el-form-item label="业务子网:"  prop="network_id" >
              <el-select
                  v-model="formData.network_id"
                  style="width:256px"
                  :disabled="isOptionIpSubnetServiceDisabled"
                  placeholder="请选择业务子网"
                  clearable
              >
                <el-option
                    v-for="item in ipSubnetServices"
                    :key="item.ID"
                    :value="item.ID"
                    :label="item.network"
                />
              </el-select>
              <el-icon
                  class="cursor-pointer ml-2 text-gray-600"
                  @click="getIdcIpSubnetServicesByIdcId"
              ><refresh /></el-icon>
              <el-icon
                  class="cursor-pointer ml-2 text-gray-600"
                  @click="goIpSubnet"
              ><document-add /></el-icon>
            </el-form-item>
            <el-form-item label="业务IP:"  prop="ip" >
              <el-input v-model="formData.ip" :clearable="false"  placeholder="请输入业务IP"  @change="handleIpService" />
            </el-form-item>
            <el-form-item label="管理网段:"  prop="segmentIdManage" >
              <el-select
                  v-model="formData.segmentIdManage"
                  style="width:256px"
                  :disabled="isOptionIpSegmentManageDisabled"
                  placeholder="请选择管理网段"
                  clearable
                  @change="getIdcIpSubnetManagesByIdcId"
              >
                <el-option
                    v-for="item in ipSegmentManages"
                    :key="item.ID"
                    :value="item.ID"
                    :label="item.name"
                />
              </el-select>
              <el-icon
                  class="cursor-pointer ml-2 text-gray-600"
                  @click="getIdcIpSegmentManagesByIdcId"
              ><refresh /></el-icon>
              <el-icon
                  class="cursor-pointer ml-2 text-gray-600"
                  @click="goIpSegment"
              ><document-add /></el-icon>
            </el-form-item>
            <el-form-item label="管理子网:"  prop="manage_network_id" >
              <el-select
                  v-model="formData.manage_network_id"
                  style="width:256px"
                  :disabled="isOptionIpSubnetManageDisabled"
                  placeholder="请选择管理子网"
                  clearable
              >
                <el-option
                    v-for="item in ipSubnetManages"
                    :key="item.ID"
                    :value="item.ID"
                    :label="item.network"
                />
              </el-select>
              <el-icon
                  class="cursor-pointer ml-2 text-gray-600"
                  @click="getIdcIpSubnetManagesByIdcId"
              ><refresh /></el-icon>
              <el-icon
                  class="cursor-pointer ml-2 text-gray-600"
                  @click="goIpSubnet"
              ><document-add /></el-icon>
            </el-form-item>
            <el-form-item label="管理IP:"  prop="manage_ip" >
              <el-input v-model="formData.manage_ip" :clearable="false"  placeholder="请输入管理IP"  @change="handleIpManage" />
            </el-form-item>
            <el-form-item label="硬件配置:"  prop="hardware_id" >
              <el-input v-model.number="formData.hardware_id" :clearable="false" placeholder="请输入硬件配置" />
            </el-form-item>
            <el-form-item label="系统:"  prop="kickstart_id" >
              <el-input v-model.number="formData.kickstart_id" :clearable="false" placeholder="请输入系统" />
            </el-form-item>
            <el-form-item label="房间:"  prop="room_id" >
              <el-select
                  v-model="formData.room_id"
                  style="width:256px"
                  :disabled="isOptionRoomDisabled"
                  placeholder="请选择房间"
                  clearable
                  @change="handleRoomChange"
              >
                <el-option
                    v-for="item in rooms"
                    :key="item.ID"
                    :value="item.ID"
                    :label="item.name"
                />
              </el-select>
              <el-icon
                  class="cursor-pointer ml-2 text-gray-600"
                  @click="getIdcRoomsByIdcId"
              ><refresh /></el-icon>
              <el-icon
                  class="cursor-pointer ml-2 text-gray-600"
                  @click="goRoom"
              ><document-add /></el-icon>
            </el-form-item>
            <el-form-item label="机柜:"  prop="cabinet_id" >
              <el-select
                  v-model="formData.cabinet_id"
                  style="width:256px"
                  :disabled="isOptionCabinetDisabled"
                  placeholder="请选择机柜"
                  clearable
              >
                <el-option
                    v-for="item in cabinets"
                    :key="item.ID"
                    :value="item.ID"
                    :label="item.name"
                />
              </el-select>
              <el-icon
                  class="cursor-pointer ml-2 text-gray-600"
                  @click="getIdcCabinetsByRoomId"
              ><refresh /></el-icon>
              <el-icon
                  class="cursor-pointer ml-2 text-gray-600"
                  @click="goCabinet"
              ><document-add /></el-icon>
            </el-form-item>
            <el-form-item label="U位:"  prop="cabinet_u" >
              <el-input v-model="formData.cabinet_u" :clearable="false"  placeholder="请输入U位" />
            </el-form-item>
            <el-form-item label="状态:"  prop="status" >
              <el-select v-model="formData.status" placeholder="请选择状态" style="width:100%" :clearable="false" >
                <el-option v-for="(item,key) in DeviceStatusOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="开发:"  prop="manager_dev" >
              <el-input v-model="formData.manager_dev" :clearable="false"  placeholder="请输入开发负责人" />
            </el-form-item>
            <el-form-item label="运维:"  prop="manager_ops" >
              <el-input v-model="formData.manager_ops" :clearable="false"  placeholder="请输入运维负责人" />
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
                <el-descriptions-item label="SN">
                        {{ formData.sn }}
                </el-descriptions-item>
                <el-descriptions-item label="资产编号">
                        {{ formData.asset_id }}
                </el-descriptions-item>
                <el-descriptions-item label="实例名">
                        {{ formData.hostname }}
                </el-descriptions-item>
                <el-descriptions-item label="业务IP">
                        {{ formData.ip }}
                </el-descriptions-item>
                <el-descriptions-item label="业务网络">
                        {{ formData.network_id }}
                </el-descriptions-item>
                <el-descriptions-item label="管理IP">
                        {{ formData.manage_ip }}
                </el-descriptions-item>
                <el-descriptions-item label="管理网络">
                        {{ formData.manage_network_id }}
                </el-descriptions-item>
                <el-descriptions-item label="硬件配置">
                        {{ formData.hardware_id }}
                </el-descriptions-item>
                <el-descriptions-item label="系统">
                        {{ formData.kickstart_id }}
                </el-descriptions-item>
                <el-descriptions-item label="机柜">
                        {{ formData.cabinet_id }}
                </el-descriptions-item>
                <el-descriptions-item label="U位">
                        {{ formData.cabinet_u }}
                </el-descriptions-item>
                <el-descriptions-item label="状态">
                        {{ filterDict(formData.status,DeviceStatusOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="开发">
                        {{ formData.manager_dev }}
                </el-descriptions-item>
                <el-descriptions-item label="运维">
                        {{ formData.manager_ops }}
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
  createDeviceBareMetal,
  deleteDeviceBareMetal,
  deleteDeviceBareMetalByIds,
  updateDeviceBareMetal,
  findDeviceBareMetal,
  getDeviceBareMetalList,
} from '@/api/deviceBareMetal'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, watch } from 'vue'

import { useRouter } from 'vue-router'
const router = useRouter()

const goIdc = () => {
  router.push({ name: 'info' })
}

const goIpSegment = () => {
  router.push({ name: 'ipSegment' })
}

const goIpSubnet = () => {
  router.push({ name: 'ipSubnet' })
}

const goRoom = () => {
  router.push({ name: 'room' })
}

const goCabinet = () => {
  router.push({ name: 'cabinet' })
}

defineOptions({
    name: 'DeviceBareMetal'
})

// 自动化生成的字典（可能为空）以及字段
const DeviceStatusOptions = ref([])
const formData = ref({
        sn: '',
        asset_id: '',
        hostname: '',
        ip: '',
        manage_ip: '',
        cabinet_u: '',
        status: undefined,
        manager_dev: '',
        manager_ops: '',
        label: '',
        describe: '',
})

// 获取未纳管BareMetalSn列表
import {
  getServerInfoSnListNotExistsInBareMetal,
} from '@/api/serverInfo'

const sns = ref([])
const getSns = async() => {
  const res = await getServerInfoSnListNotExistsInBareMetal()
  if (res.code === 0) {
    sns.value = res.data.list
  }
}

// 获取IDC列表
import {
  getIdcInfoList
} from '@/api/idcInfo'

const idcs = ref([])
const getIdcs = async() => {
  const res = await getIdcInfoList()
  if (res.code === 0) {
    idcs.value = res.data.list
  }
}

const handleIdcChange = () => {
  getIdcRoomsByIdcId();
  getIdcIpSegmentServicesByIdcId();
  getIdcIpSegmentManagesByIdcId();
  getSns();
}

// 根据IdcId获取IpSegment列表
import {
  getIdcIpSegmentList
} from '@/api/idcIpSegment'

const ipSegmentServices = ref([])
const isOptionIpSegmentServiceDisabled = ref(true);
const getIdcIpSegmentServicesByIdcId = () => {
  // 调用网段列表接口，并传递IdcId的值作为参数
  const { idc_id } = formData.value;
  getIdcIpSegmentList({ IdcId: idc_id, ipSegmentType: 0 }).then(response => {
    // 根据接口返回的数据更新网段列表
    ipSegmentServices.value = response.data.list;
    isOptionIpSegmentServiceDisabled.value = idc_id === 0;
  }).catch(error => {
    console.error("获取业务网段列表失败:", error);
  });
};

const ipSegmentManages = ref([])
const isOptionIpSegmentManageDisabled = ref(true);
const getIdcIpSegmentManagesByIdcId = () => {
  // 调用网段列表接口，并传递IdcId的值作为参数
  const { idc_id } = formData.value;
  getIdcIpSegmentList({ IdcId: idc_id ,ipSegmentType: 1}).then(response => {
    // 根据接口返回的数据更新网段列表
    ipSegmentManages.value = response.data.list;
    isOptionIpSegmentManageDisabled.value = idc_id === 0;
  }).catch(error => {
    console.error("获取管理网段列表失败:", error);
  });
};

// 根据SegmentId获取IpSubnet列表
import {
  getIdcIpSubnetList,validateIpSubnet
} from '@/api/idcIpSubnet'

const ipSubnetServices = ref([])
const isOptionIpSubnetServiceDisabled = ref(true);
const getIdcIpSubnetServicesByIdcId = () => {
  // 调用B接口，并传递选项A的值作为参数
  const { segmentIdService } = formData.value;
  getIdcIpSubnetList({ SegmentId: segmentIdService}).then(response => {
    // 根据接口返回的数据更新选项B的列表
    ipSubnetServices.value = response.data.list;
    isOptionIpSubnetServiceDisabled.value = segmentIdService === 0;
  }).catch(error => {
    console.error("获取业务子网列表失败:", error);
  });
};

const ipSubnetManages = ref([])
const isOptionIpSubnetManageDisabled = ref(true);
const getIdcIpSubnetManagesByIdcId = () => {
  // 调用B接口，并传递选项A的值作为参数
  const { segmentIdManage } = formData.value;
  getIdcIpSubnetList({ SegmentId: segmentIdManage}).then(response => {
    // 根据接口返回的数据更新选项B的列表
    ipSubnetManages.value = response.data.list;
    isOptionIpSubnetManageDisabled.value = segmentIdManage === 0;
  }).catch(error => {
    console.error("获取管理子网列表失败:", error);
  });
};

const messageIpService = ref("");
const handleIpService= () => {
  const { ip, network_id } = formData.value;
  validateIpSubnet({ ID: network_id, address: ip }).then(response => {
    // 根据接口返回的数据更新选项B的列表
    messageIpService.value = response.Message
  }).catch(error => {
    messageIpService.value = "业务IP校验失败"
  });
};

const messageIpManage = ref("");
const handleIpManage = () => {
  const { manage_ip, manage_network_id } = formData.value;
  validateIpSubnet({ ID: manage_network_id, address: manage_ip }).then(response => {
    // 根据接口返回的数据更新选项B的列表
    messageIpManage.value = response.Message
  }).catch(error => {
    messageIpManage.value = "管理IP校验失败"
  });
};
// 根据IdcId获取Room列表
import {
  getIdcRoomList
} from '@/api/idcRoom'

const rooms = ref([])
const isOptionRoomDisabled = ref(true);
const getIdcRoomsByIdcId = () => {
  // 调用B接口，并传递选项A的值作为参数
  // 假设接口调用方法为callB接口，且需要传递参数optionA
  const { idc_id } = formData.value;
  getIdcRoomList({ IdcId: idc_id }).then(response => {
    // 根据接口返回的数据更新选项B的列表
    rooms.value = response.data.list;
    isOptionRoomDisabled.value = idc_id === 0;
  }).catch(error => {
    console.error("获取房间列表失败:", error);
  });
};

const handleRoomChange = () => {
  getIdcCabinetsByRoomId();
}
// 根据RoomId获取Cabinet列表
import {
  getIdcCabinetList
} from '@/api/idcCabinet'
import {entryServerDiscoveryByIds} from "@/api/serverDiscovery";

const cabinets = ref([])
const isOptionCabinetDisabled = ref(true);
const getIdcCabinetsByRoomId = () => {
  // 调用B接口，并传递选项A的值作为参数
  // 假设接口调用方法为callB接口，且需要传递参数optionA
  const { room_id, cabinet_id } = formData.value;
  getIdcCabinetList({ RoomId: room_id }).then(response => {
    // 根据接口返回的数据更新选项B的列表
    cabinets.value = response.data.list;
    isOptionCabinetDisabled.value = room_id === 0;
  }).catch(error => {
    console.error("获取机柜列表失败:", error);
  });
};

// 验证规则
const rule = reactive({
               sn : [{
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
               asset_id : [{
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
               hostname : [{
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
               ip : [{
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
               network_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               manage_ip : [{
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
               manage_network_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               kickstart_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               cabinet_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               cabinet_u : [{
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
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               manager_dev : [{
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
               manager_ops : [{
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
  const table = await getDeviceBareMetalList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    DeviceStatusOptions.value = await getDictFunc('DeviceStatus')
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
            deleteDeviceBareMetalFunc(row)
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
      const res = await deleteDeviceBareMetalByIds({ ids })
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
const updateDeviceBareMetalFunc = async(row) => {
    const res = await findDeviceBareMetal({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.redeviceBareMetal
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteDeviceBareMetalFunc = async (row) => {
    const res = await deleteDeviceBareMetal({ ID: row.ID })
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
  const res = await findDeviceBareMetal({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.redeviceBareMetal
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          sn: '',
          asset_id: '',
          hostname: '',
          ip: '',
          manage_ip: '',
          cabinet_u: '',
          status: undefined,
          manager_dev: '',
          manager_ops: '',
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
        sn: '',
        asset_id: '',
        hostname: '',
        ip: '',
        manage_ip: '',
        cabinet_u: '',
        status: undefined,
        manager_dev: '',
        manager_ops: '',
        label: '',
        describe: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

import {
 osInstallStart
} from '@/api/osInstallQueue'

const osInstall = async(row) => {
  const res = await osInstallStart(row)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '开始重装'
    })
    getTableData()
  }
}

const init = () => {
  getIdcs()
}

init()

</script>

<style>

</style>
