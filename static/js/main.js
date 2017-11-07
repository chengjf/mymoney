axios.defaults.baseURL = 'http://127.0.0.1:8088';
axios.defaults.headers.common['Authorization'] = "abcdefghijl";
axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';
Vue.prototype.$http = axios

var app = new Vue({

    el: '#app',
    data: {
        message: 'Hello Vue!',
        debitAccount: "",
        debitEntry: "",
        creditAccount: "",
        creditEntry: "",
        amount: 0,
        datetime: '',
        counter: '无',
        accounts: [],
        entries: [],

        //
        firstEntry: "",
        secondEntry: "",
        thirdEntry: "",
        firstEntries: [],
        secondEntries: [],
        thirdEntries: [],
        date: "",
        time: ""
    },
    computed: {},
    mounted: function () {
        this.$nextTick(function () {
            that = this
            this.$http.get('/accounts').then(function (res) {
                console.log(res.data)
                that.accounts = res.data;
            })
            this.$http.get('/entries').then(function (res) {
                console.log("first entries: ", res.data)
                that.firstEntries = res.data;
            })
        })
    },
    watch: {
        /**
         * On update of options-prop, recreate element
         */
        // 'options': function() {
        //     console.log("options watch")
        //     this.create();
        // },
        options: {
            handler: function (val, oldVal) {
                this.create(); // call it in the context of your component object
            },
            deep: true
        }
    },
    methods: {

        clear: function () {
            var el = this.$el;
            $(el).select2('val', '');
        },
        create: function () {
            // var el = this.$el;
            // $(el).select2().on('change', function(e) {
            //     this.$set('val', $(el).select2('data').map(function(d) { return { value: d.id, text: d.text } }));
            // }.bind(this));
            console.log("create")
        },
        submit: function () {
            const postData = {
                debitAccount: this.debitAccount,
                debitEntry: this.debitEntry,
                creditAccount: this.creditAccount,
                creditEntry: this.creditEntry,
                amount: Number(this.amount),
                datetime: this.date + " " + this.time+":00",
                counter: this.counter,
            }
            this.$http.post("/createRecord", postData).then(function (res) {
                if (res.status === 200) {
                    console.log(res.data)
                    alert("保存成功")
                } else {
                    alert("保存失败")
                }
            }).catch(function (response) {
                console.log(response);
                alert("保存失败")
            });
            ;
        },
        changeEntry: function(name){
            console.log(name)
            that = this
            if(name === 2){
                this.$http.get("/entries/"+this.firstEntry).then(function (res) {
                    if (res.status === 200) {
                        console.log(res.data)
                        that.secondEntries = res.data
                    } else {
                    }
                }).catch(function (response) {
                    console.log(response);
                });
            }else if(name == 3){
                this.$http.get("/entries/"+this.secondEntry).then(function (res) {
                    if (res.status === 200) {
                        console.log(res.data)
                        that.thirdEntries = res.data

                    } else {
                    }
                }).catch(function (response) {
                    console.log(response);
                });
            }
            console.log("changeEntry ")
        }
    }
})
