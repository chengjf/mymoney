axios.defaults.baseURL = 'http://127.0.0.1:8080';
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
                console.log(res.data)
                that.entries = res.data;
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
        }
    }
})