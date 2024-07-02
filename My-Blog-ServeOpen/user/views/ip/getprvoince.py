from luntan.tools.ip2region.xdbSearcher import XdbSearcher

def get_province(ip):
    dbPath = "../luntan/luntan/tools/ip2region/data/ip2region.xdb"
    searcher = XdbSearcher(dbfile=dbPath)
    region_str = searcher.searchByIPStr(ip)
    searcher.close()
    province = region_str.split('|')
    return province

# print(get_province('112.26.68.69'))