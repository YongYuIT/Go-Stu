// Pitaya_CPP_Client.cpp : 此文件包含 "main" 函数。程序执行将在此处开始并结束。
//

#include <iostream>
#include <pitaya.h>
#include <cstdarg>
using namespace std;


string emptyStr = "";

void ev_handler(pc_client_t* client, int ev_type, void* ex_data, const char* arg1, const char* arg2) {

	cout << "ev_handler-->ex_data::";
	if (ex_data) {
		char* ext_str = (char*)ex_data;
		cout << ext_str;
	}
	if (arg1) {
		cout << "-->arg1::" << arg1;
	}
	if (arg2) {
		cout << "-->arg2::" << arg2;
	}
	cout << endl;



}

void push_handler(pc_client_t* client, const char* route, const pc_buf_t* payload) {
	cout << "push_handler-->route::" << route << endl;
}

void quiet_log(int level, const char* msg, ...)
{
	cout << "Log Level: " << level << " Message: ";
	va_list args;
	va_start(args, msg);
	vfprintf(stdout, msg, args); // 使用 vfprintf 进行格式化输出
	va_end(args);
	cout << endl;
}

void request_success_handler(const pc_request_t* req, const pc_buf_t* resp) {
	std::string buff((char*)resp->base, resp->len);
	cout << "request success-->" << buff << endl;
}

void request_error_handler(const pc_request_t* req, const pc_error_t* error) {
	char* ext_str = (char*)pc_request_ex_data(req);
	cout << "request failed!-->" << ext_str << endl;
}


int main()
{
	string continueStr;

	cout << "start client test" << endl;

	cout << "start lib init" << endl;
	pc_lib_client_info_t client_info;
	client_info.platform = "windows";
	client_info.build_number = "v1";
	client_info.version = "1.0";
	pc_lib_init(quiet_log, NULL, NULL, NULL, client_info);

	cout << "start client init" << endl;
	pc_client_config_t config = PC_CLIENT_CONFIG_DEFAULT;
	config.conn_timeout = 10;
	string init_tag = "hello_test_client_init";
	auto result = pc_client_init((void*)init_tag.c_str(), &config);

	cout << "start client create" << endl;
	pc_client_t* client = result.client;
	if (!client) {
		cout << "client create failed!!" << endl;
		return -1;
	}
	cout << "client create success" << endl;

	cout << "start client conn" << endl;
	string add_ev_tag = "hello_test_add_ev";
	int ev_handler_id = pc_client_add_ev_handler(client, ev_handler, (void*)add_ev_tag.c_str(), NULL);
	pc_client_set_push_handler(client, push_handler);
	string host = "localhost";
	int conn_state = pc_client_connect(client, host.c_str(), 3250, NULL);
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
	//注意，此处立即打印 "conn success"，并不可信。因为pc_client_connect需要等到10秒才会有结果
	cin >> continueStr;

	cout << "start send with no params" << endl;
	string req_tag = "hello_test_client_request_no_params";
	pc_string_request_with_timeout(client, "chat.room.join", NULL, (void*)req_tag.c_str(), 10, request_success_handler, request_error_handler);


	cout << "start send msg" << endl;
	string req_tag1 = "hello_test_client_request_with_params";
	string request_params = "{\"msg\":\"cppttttest\"}";
	pc_binary_request_with_timeout(client, "chat.room.testmsg", (uint8_t*)request_params.data(), request_params.size(), (void*)req_tag1.c_str(), 10, request_success_handler, request_error_handler);

	cout << "start send rpc msg" << endl;
	string req_tag2 = "hello_test_client_request_with_rpc_params";
	string request_params1 = "{\"msg\":\"cpprpcttttest\"}";
	pc_binary_request_with_timeout(client, "chat.room.rpctestmsg", (uint8_t*)request_params1.data(), request_params1.size(), (void*)req_tag2.c_str(), 10, request_success_handler, request_error_handler);

	cin >> continueStr;
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
