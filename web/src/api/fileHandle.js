import service from '@/utils/request'

export const resumeTransChunkContinue = (data) => {
  return service({
    url: '/fileResumeTrans/chunkContinue',
    method: 'post',
    donNotShowLoading: true,
    headers: { 'Content-Type': 'multipart/form-data' },
    data
  })
}

export const resumeTransChunkFinish = (params) => {
  return service({
    url: '/fileResumeTrans/chunkFinish',
    method: 'post',
    params
  })
}

export const resumeTransChunkRemove = (data, params) => {
  return service({
    url: '/fileResumeTrans/chunkRemove',
    method: 'post',
    data,
    params
  })
}

export const resumeTransFindFile = (params) => {
  return service({
    url: '/fileResumeTrans/findFile',
    method: 'get',
    params
  })
}