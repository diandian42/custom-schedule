// 请求工具函数（待实现）
export const request = (options) => {
  return new Promise((resolve, reject) => {
    uni.request({
      url: options.url,
      method: options.method || 'GET',
      data: options.data || {},
      header: options.header || {
        'Content-Type': 'application/json',
      },
      success: (res) => {
        resolve(res.data)
      },
      fail: (err) => {
        reject(err)
      },
    })
  })
}

