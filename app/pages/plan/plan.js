//index.js
//获取应用实例

// tab页 demo: https://www.jianshu.com/p/1a3405f77654

const app = getApp()

var order = ['red', 'yellow', 'blue', 'green', 'red']

Page({
  data: {
    motto: 'Hello World',
    userInfo: {},
    hasUserInfo: false,
    canIUse: wx.canIUse('button.open-type.getUserInfo'),
    winWidth: 0,
    winHeight: 0,
    currentTab: 0,
    toView: 'green',
    scrollTop: 100,
    scrollLeft: 0,
    datas: [
      "https://gitee.com/index/ent_poster/banner_5_1_a.png",
      "https://gitee.com/index/ent_poster/banner_5_1_a.png",
      "https://gitee.com/index/ent_poster/banner_5_1_a.png",
      "https://gitee.com/index/ent_poster/banner_5_1_a.png",
      "https://gitee.com/index/ent_poster/banner_5_1_a.png",
      "https://gitee.com/index/ent_poster/banner_5_1_a.png",
      "https://gitee.com/index/ent_poster/banner_5_1_a.png",
      "https://gitee.com/index/ent_poster/banner_5_1_a.png",
      "https://gitee.com/index/ent_poster/banner_5_1_a.png",
      "https://gitee.com/index/ent_poster/banner_5_1_a.png",
      "https://gitee.com/index/ent_poster/banner_5_1_a.png",
      "https://gitee.com/index/ent_poster/banner_5_1_a.png",            
      "https://gitee.com/index/ent_poster/banner_5_1_a.png"
    ]
  },

  a: function () {
   
  },

  onLoad: function() {

    var that = this;

    /**
     * 获取当前设备的宽高
     */
    wx.getSystemInfo( {

        success: function( res ) {
            that.setData( {
                winWidth: res.windowWidth,
                winHeight: res.windowHeight
            });
        }

    });
  },
    
  //  tab切换逻辑
  swichNav: function( e ) {

      var that = this;

      if( this.data.currentTab === e.target.dataset.current ) {
          return false;
      } else {
          that.setData( {
              currentTab: e.target.dataset.current
          })
      }
  },

  bindChange: function( e ) {

      var that = this;
      that.setData( { currentTab: e.detail.current });

  },


  //事件处理函数
  bindViewTap: function() {
    wx.navigateTo({
      url: '../logs/logs'
    })
  },

  //滚动条滚到顶部的时候触发
  upper: function(e) {
    console.log("顶部");
  },
  //滚动条滚到底部的时候触发
  lower: function(e) {
    console.log("底");
  },
  //滚动条滚动后触发
  scroll: function(e) {
    console.log("滚动");
  },
  //通过设置滚动条位置实现画面滚动
  tapMove: function(e) {
    this.setData({
      scrollTop: this.data.scrollTop + 10
    })
  }


})
