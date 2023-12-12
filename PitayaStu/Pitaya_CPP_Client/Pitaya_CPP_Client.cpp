// Pitaya_CPP_Client.cpp : 此文件包含 "main" 函数。程序执行将在此处开始并结束。
//

#include <iostream>
#include <pitaya.h>
using namespace std;

void ev_handler(pc_client_t* client, int ev_type, void* ex_data, const char* arg1, const char* arg2) {

}

void push_handler(pc_client_t* client, const char* route, const pc_buf_t* payload) {

}

void quiet_log(int level, const char* msg, ...)
{
	cout << "level::" << level << "-->msg::" << msg << endl;
}

int main()
{
	cout << "start client test" << endl;

	cout << "start lib init" << endl;
	pc_lib_client_info_t client_info;
	client_info.platform = "windows";
	client_info.build_number = "v1";
	client_info.version = "1.0";
	pc_lib_init(quiet_log, NULL, NULL, NULL, client_info);

	cout << "start client init" << endl;
	pc_client_config_t config = PC_CLIENT_CONFIG_DEFAULT;
	config.conn_timeout = 30;
	auto result = pc_client_init(NULL, &config);

	cout << "start client create" << endl;
	pc_client_t* client = result.client;
	if (!client) {
		cout << "client create failed!!" << endl;
		return -1;
	}
	cout << "client create sucess" << endl;

	cout << "start client conn" << endl;
	int ev_handler_id = pc_client_add_ev_handler(client, ev_handler, NULL, NULL);
	pc_client_set_push_handler(client, push_handler);
	int conn_state = pc_client_connect(client, "localhost", 3250, NULL);
	switch (conn_state)
	{
	case PC_RC_OK:
		cout << "conn success" << endl;
		break;
	case PC_RC_INVALID_JSON:
		cout << "conn failed for invalid json" << endl;
		break;
	default:
		cout << "conn failed" << endl;
	}

}

// 运行程序: Ctrl + F5 或调试 >“开始执行(不调试)”菜单
// 调试程序: F5 或调试 >“开始调试”菜单

// 入门使用技巧: 
//   1. 使用解决方案资源管理器窗口添加/管理文件
//   2. 使用团队资源管理器窗口连接到源代码管理
//   3. 使用输出窗口查看生成输出和其他消息
//   4. 使用错误列表窗口查看错误
//   5. 转到“项目”>“添加新项”以创建新的代码文件，或转到“项目”>“添加现有项”以将现有代码文件添加到项目
//   6. 将来，若要再次打开此项目，请转到“文件”>“打开”>“项目”并选择 .sln 文件
