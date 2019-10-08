#### type

    int     默认 int8 -> int16, int32, int64
    uint    默认 uint8, 超出自动申请更高空间 -> uint16, uint32, uint64
    double  默认 float64
    string  默认 [255]byte -> [512]byte -> ...
    boolean 默认 bool
    array   默认 struct{}, []interface{}
    object  默认 struct{}
    ptr     默认 ptr
    *       指针, 取值
    &       取址

#### variable

    int i = 1
    double db = 1.0
    string a = "123456"
    boolean b = false
    array arr = ["1", "2", "3", 4, 7]
    object obj = {}
    ptr ii = &i
    int i2 = *ii
    *ii = 2


#### function 

    function functionName (type variable1) (type){
        return 0
    }

#### for

    for i=0;i<10;i++ {

    }
