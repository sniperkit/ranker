package main

var (
	cityMap = map[string]string{
		"shang,hai,,china":                                                                                      "shanghai",
		"dalian,,china":                                                                                         "liaoning",
		"nanning,,guangxi,,china":                                                                               "guangxi",
		"harbin,heilongjiang,china":                                                                             "heilongjiang",
		"yangling,,shaanxi,,china":                                                                              "shaanxi",
		"china,beijing.":                                                                                        "beijing",
		"\u5317\u4eac-\u671d\u9633\u533a(china-beijing)":                                                        "beijing",
		"jiangmen-guangdong-china":                                                                              "guangdong",
		"hangchow,,china":                                                                                       "zhejiang",
		"china,zhejiang,hangzhou":                                                                               "zhejiang",
		"wuhan,,china":                                                                                          "hubei",
		"qingdao,,china":                                                                                        "shandong",
		"shangdi,,haidian,district,,beijing,100085,china":                                                       "beijing",
		"shenzhen,,china,&,montreal,,canada":                                                                    "guangdong",
		"fuzhou,,china":                                                                                         "fujian",
		"dalian,,liaoning,,china":                                                                               "liaoning",
		"tianjin/beijing,,china":                                                                                "tianjin",
		"qingdao,,shandong,,china":                                                                              "shandong",
		"800,dongchuan,rd.,minhang,district,,shanghai,,china":                                                   "shanghai",
		"berlin,,germany;,porto-novo,,benin;,accra,,ghana;,nairobi,,kenya;,beijing,,china;,africa":              "beijing",
		"china,guangdong,(\u5e7f\u5dde)":                                                                        "guangdong",
		"huizhou,,china":                                                                                        "guangdong",
		"wuhan,china":                                                                                           "hubei",
		"changsha,china":                                                                                        "hunan",
		"china,chengdu":                                                                                         "sichuan",
		"b\u011bij\u012bng,,china":                                                                              "beijing",
		"shanghai,,china.":                                                                                      "shanghai",
		"ningbo,,zhejiang,,china":                                                                               "zhejiang",
		"wuxi,china":                                                                                            "jiangsu",
		"shenzhen,,guangdong,,china":                                                                            "guangdong",
		"china,changchun":                                                                                       "jilin",
		"jiangmen,guangdong,china":                                                                              "guangdong",
		"china,,\u82cf\u5dde":                                                                                   "jiangsu",
		"fuzhou,,fujian,,china":                                                                                 "fujian",
		"xiamen,china":                                                                                          "fujian",
		"china-guangdong":                                                                                       "guangdong",
		"beijing,,,china":                                                                                       "beijing",
		"china,,shanghai~":                                                                                      "shanghai",
		"tianjin,,china":                                                                                        "tianjin",
		"shanghai,.,china":                                                                                      "shanghai",
		"dali,,china":                                                                                           "yunnan",
		"peking,china":                                                                                          "beijing",
		"shanghai,,mainland,china":                                                                              "shanghai",
		"canton,,china":                                                                                         "guangdong",
		"hangzhou,,shanghai,china":                                                                              "shanghai",
		"beijing,,p.,r.,china":                                                                                  "beijing",
		"china-fujian-fuzhou":                                                                                   "fujian",
		"shantou,,,guangdong,,,china":                                                                           "guangdong",
		"dongguan,,china":                                                                                       "guangdong",
		"melbourne,,australia,(shanghai,,china)":                                                                "shanghai",
		"hubei,,china":                                                                                          "hubei",
		"china,guangdong,shenzheng":                                                                             "guangdong",
		"zhengzhou,henan,,china":                                                                                "henan",
		"amoy,,china":                                                                                           "fujian",
		"changsha,,hunan,,china":                                                                                "hunan",
		"china,wuhan":                                                                                           "hubei",
		"harbin,,heilongjiang,,china":                                                                           "heilongjiang",
		"changde,hunan,china":                                                                                   "hunan",
		"q\u012bngd\u01ceo,china":                                                                               "shandong",
		"china/hangzhou":                                                                                        "zhejiang",
		"\bbeijing,,china":                                                                                      "beijing",
		"chengdu,in,sichuan,of,china":                                                                           "sichuan",
		"beijing,.,china":                                                                                       "beijing",
		"wenzhou,china":                                                                                         "zhejiang",
		"jinan,,china":                                                                                          "shandong",
		"xi'an,,shaanxi,,china":                                                                                 "shaanxi",
		"guangzhou,guangdong,china":                                                                             "guangdong",
		"sichuan,,chengdu,,china":                                                                               "sichuan",
		"china,,,beijing":                                                                                       "beijing",
		"guizhou,china":                                                                                         "hubei",
		"hangzhou,,china.":                                                                                      "zhejiang",
		"\u4e0a\u6d77,china":                                                                                    "shanghai",
		"huizhou,china":                                                                                         "guangdong",
		"shanghai,in,china":                                                                                     "shanghai",
		"zhengzhou,henan,china":                                                                                 "henan",
		"mount,wuyi,of,nanping,city,,fujian,province,,china":                                                    "fujian",
		"unit,1801,,tower,1,,china,central,place,office,building,,no.81,jianguo,rd,,chaoyang,district,,beijing": "beijing",
		"shanghai,(china)":                        "shanghai",
		"beijing,china.":                          "beijing",
		"chaoyang,,china":                         "liaoning",
		"harbin.,china":                           "heilongjiang",
		"xiaoshan,,zhejiang,,china":               "zhejiang",
		"ningbo,,china":                           "zhejiang",
		"chengdu,&,langfang,,china":               "sichuan",
		"china,nanjing":                           "jiangsu",
		"china,,wuhan":                            "hubei",
		"gz,,china":                               "guangdong",
		"shenzhen,china":                          "guangdong",
		"china,jiangsu,nanjing":                   "jiangsu",
		"hangzhou,china":                          "zhejiang",
		"beijing,,mainland,china":                 "beijing",
		"shenzhen,,china":                         "guangdong",
		"da,lian,china":                           "liaoning",
		"zhenjiang,city,,jiangsu,province,,china": "jiangsu",
		"hangzhou,/,shanghai,,china":              "shanghai",
		"lianyungang,,china":                      "jiangsu",
		"shanghai,,china":                         "shanghai",
		"xian,china":                              "shaanxi",
		"chengdu,,sichuan,,china":                 "sichuan",
		"west,lake,,,hangzhou,,,china":            "zhejiang",
		"henan,province,,china":                   "henan",
		"chengdu,china":                           "sichuan",
		"nanking,,china":                          "jiangsu",
		"hangzhou\uff0czhejiang\uff0cchina":       "zhejiang",
		"china-shenzhen":                          "guangdong",
		"shandong,,china":                         "shandong",
		"china,hangzhou":                          "zhejiang",
		"fuzhou,china":                            "fujian",
		"china,,beijing":                          "beijing",
		"china,,shang,hai":                        "shanghai",
		"beijing,of,china":                        "beijing",
		"zhuhai,,china":                           "guangdong",
		"chengdu,sichuan,china":                   "sichuan",
		"guangzhou,,guangdong,,china/hk":          "guangdong",
		"beijing,/,shenzhen,,china":               "beijing",
		"pudong,district,,shanghai,,china":        "shanghai",
		"china,guangdong,huizhou":                 "guangdong",
		"shenyang,china":                          "liaoning",
		"zhenjiang,,china":                        "jiangsu",
		"peking,,china":                           "beijing",
		"china,,guangzhou":                        "guangdong",
		"quanzhou,,china":                         "fujian",
		"shanghai,china":                          "shanghai",
		"china,,shaanxi":                          "shaanxi",
		"chengdu,of,china":                        "sichuan",
		"kunming,,yunnan,,china":                  "yunnan",
		"china,,hunan,,changsha":                  "hunan",
		"wuhan,,hubei,china":                      "hubei",
		"china,hubei,wuhan":                       "hubei",
		"beijing\uff0cchina":                      "beijing",
		"china,\u6c5f\u82cf":                      "jiangsu",
		"beijing,,china.":                         "beijing",
		"xiamen,fujian,china":                     "fujian",
		"changchun,,china":                        "jilin",
		"bejing,,china":                           "beijing",
		"chongqing,,china":                        "chongqing",
		"china-\u5e7f\u5dde":                      "guangdong",
		"wuxi,,jiangsu,,china":                    "jiangsu",
		"xi'an/beijing,,china":                    "beijing",
		"college,park,,md;,beijing,,china":        "beijing",
		"p.r.,china":                              "beijing",
		"vancouver,,canada;,beijing,,china":       "beijing",
		"china-chengdu":                           "sichuan",
		"nanning,,china":                          "guangxi",
		"xia,men,fu,jian,china":                   "fujian",
		"hang,zhou,,china":                        "zhejiang",
		"shenzhen,guangdong,china":                "guangdong",
		"nanjing,,jiangsu,province,,china":        "jiangsu",
		"beijing,,china,/,apeldoorn,,nl":          "beijing",
		"china,\u82cf\u5dde":                      "jiangsu",
		"changsha,hunan,china":                    "hunan",
		"hefei,,china":                            "anhui",
		"shanghai.,china":                         "shanghai",
		"china/guangdong/guangzhou":               "guangdong",
		"zhejiang,china":                          "zhejiang",
		"shanghai,,,china":                        "shanghai",
		"wuhan,,hubei,,china":                     "hubei",
		"zhengzhou,,henan,,china":                 "henan",
		"hefei,china":                             "anhui",
		"haerbin,,china":                          "heilongjiang",
		"guangzhou,china":                         "guangdong",
		"shaanxi,china":                           "shaanxi",
		"china,shanghai":                          "shanghai",
		"china,,hangzhou":                         "zhejiang",
		"guang,dong,,china":                       "guangdong",
		"xi'an,china":                             "shaanxi",
		"nyc(beijing,,china)":                     "beijing",
		"tangshan,,china":                         "hebei",
		"beijing,,china":                          "beijing",
		"quanzhou,,fujian,,china":                 "fujian",
		"zhuhai,china":                            "guangdong",
		"shenyang,,china":                         "liaoning",
		"jiangsu,university,,china":               "jiangsu",
		"nanjing,,china":                          "jiangsu",
		"\bshenzhen,,china":                       "guangdong",
		"sichuan,china":                           "sichuan",
		"a207-1,,456,bibo,road,,zhang,jiang,high-tech,park,,shanghai,china": "shanghai",
		"shanhai,china":                            "shanghai",
		"beihang,university,,beijing,,china":       "beijing",
		"huadu,district,guangzhou,guangdong,china": "guangdong",
		"beijing,,beijing,,china":                  "beijing",
		"changsha,,china":                          "hunan",
		"chengdu,university,of,information,technology,chengdu,sichuan,china.": "sichuan",
		"guangzhou,,guangdong,,china":                                         "guangdong",
		"pudong,,shanghai,,china":                                             "shanghai",
		"zhongshan,,china":                                                    "guangdong",
		"dalian,liaoning,china":                                               "liaoning",
		"nanjing,city,,,china":                                                "fujian",
		"qihoo,360,,beijing,,china":                                           "beijing",
		"shanghai,,,republic,of,china":                                        "shanghai",
		"shanghai,,shanghai,,china":                                           "shanghai",
		"zaozhuang,china":                                                     "shandong",
		"china,,guangdong,,guangzhou,,jinan,university":                       "guangdong",
		"china,canton,guangzhou":                                              "guangdong",
		"fuzhou,fujian,china":                                                 "fujian",
		"sh,,china":                                                           "hebei",
		"china,hanzhou":                                                       "yunnan",
		"haerbin,china":                                                       "heilongjiang",
		"china,shenzhen":                                                      "guangdong",
		"sunnyvale,,ca;,beijing,,china":                                       "beijing",
		"kunshan,china":                                                       "jiangsu",
		"harbin,,china":                                                       "heilongjiang",
		"changchun,,jilin,,china":                                             "jilin",
		"hangzhou,,zhejiang,,china":                                           "zhejiang",
		"china,canton,province":                                               "guangdong",
		"beijing,.china":                                                      "beijing",
		"wuhan,hubei,china":                                                   "hubei",
		"china,sichuan,chengdu":                                               "sichuan",
		"\u6c5f\u82cf\u5f90\u5dde@shanghai,china":                             "shanghai",
		"minhang,,china":                                                      "shanghai",
		"hz.,china":                                                           "gansu",
		"shanghai,,ningbo\uff0c,china":                                        "shanghai",
		"china,beijing":                                                       "beijing",
		"xi'an,,china":                                                        "shaanxi",
		"shanghai,,china,,earth":                                              "shanghai",
		"yueyang,,china":                                                      "hunan",
		"zhengzhou,,china":                                                    "henan",
		"chongqing,china":                                                     "chongqing",
		"yantai,,shandong,,china":                                             "shandong",
		"china,hangzhou,\u897f\u6e56\u533a":                                   "zhejiang",
		"china,beijing,tongzhou":                                              "beijing",
		"xiamen,,china":                                                       "fujian",
		"china,linyi":                                                         "shandong",
		"zhejiang,hangzhou,,china":                                            "zhejiang",
		"bejing,china":                                                        "beijing",
		"beijing,china":                                                       "beijing",
		"canton,china":                                                        "guangdong",
		"suzhou,china":                                                        "jiangsu",
		"suzhou,,jiangsu,,china":                                              "jiangsu",
		"china,dalian":                                                        "liaoning",
		"xiamen,[amoy],,china":                                                "fujian",
		"beijing,-,china":                                                     "beijing",
		"shanghai,p.,r.,china":                                                "shanghai",
		"china,,gui,lin":                                                      "guangxi",
		"shanghai,,china,/,tokyo,,japan":                                      "shanghai",
		"\u2708,china,beijing":                                                "beijing",
		"linyi,shandong,,china":                                               "shandong",
		"foshan,china":                                                        "guangdong",
		"guangdong,,china":                                                    "guangdong",
		"china,zhejiang":                                                      "zhejiang",
		"hangzhou,zhejiang,china":                                             "zhejiang",
		"hangzhou.,china":                                                     "zhejiang",
		"chengdu,,china":                                                      "sichuan",
		"lanzhou,,china":                                                      "gansu",
		"china,huizhou,city,,guangdong":                                       "guangdong",
		"wuhan,,,hubei,,,china":                                               "hubei",
		"chengdu,,sichuan,,china.":                                            "sichuan",
		"beijing/china":                                                       "beijing",
		"suzhou,jiangsu,china":                                                "jiangsu",
		"chaozhou,,guangdong,,china":                                          "guangdong",
		"zhengzhou,china":                                                     "henan",
		"\u5317\u4eac,beijing,,china":                                         "beijing",
		"beijing,,china,/,apeldoorn,,netherlands":                             "beijing",
		"nanjing,china":                                                       "jiangsu",
		"qinhuangdao,china":                                                   "hebei",
		"chu,,china":                                                          "henan",
		"shenzhen,city,,guangdong,,china.":                                    "guangdong",
		"lixiang,intl,plaza,beijing,china":                                    "beijing",
		"xi'an,of,china":                                                      "shaanxi",
		"high-tech,park,shenzhen,city,guangdong,province,,china":              "guangdong",
		"guangzhou,,china":                                                    "guangdong",
		"china,sichuan":                                                       "sichuan",
		"anhui,,china":                                                        "anhui",
		"china,\u73e0\u6d77":                                                  "guangdong",
		"hefei,,anhui,,china":                                                 "anhui",
		"suzhou,,china":                                                       "jiangsu",
		"\u676d\u5dde(hangzhou),china":                                        "zhejiang",
		"china,shenyang":                                                      "liaoning",
		"xiamen,,fujian,,china":                                               "fujian",
		"nanjing,,jiangsu,,china":                                             "jiangsu",
		"\bhangzhou,,china":                                                   "zhejiang",
		"ningbo,china":                                                        "zhejiang",
		"kunming,,china":                                                      "yunnan",
		"fujian,china":                                                        "fujian",
		"hangzhou,,china":                                                     "zhejiang",
		"gz,china":                                                            "guangdong",
		"china,guangzhou":                                                     "guangdong",
		"xi'an,shaanxi,china":                                                 "shaanxi",
		"haidian,,beijing,,china":                                             "beijing",
		"china,,shenzhen":                                                     "guangdong",
		"california,,united,states;,beijing,,china":                           "beijing",
		"wuxi,,china":                                                         "jiangsu",
		"guang,zhou,,china":                                                   "guangdong",
		"shanghai,,china,and,paris,,france":                                   "shanghai",
		"china,hebei,baoding":                                                 "hebei",
		"\u4e2d\u56fd\u5317\u4eac,beijing,china":                              "beijing",
		"china,,shanghai":                                                     "shanghai",
		"yiwu,,china":                                                         "zhejiang",
		"shantou,,guangdong,,china":                                           "guangdong",
		"harbin,china":                                                        "heilongjiang",
		"maanshan,anhui,china":                                                "anhui",
		"united,states;,beijing,,china":                                       "beijing",
		"hunan,,china":                                                        "hunan",
		"ningbo->shanghai,,china":                                             "shanghai",
		"china::jinan":                                                        "shandong",
	}
)
