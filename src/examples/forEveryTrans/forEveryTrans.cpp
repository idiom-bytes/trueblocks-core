/*-------------------------------------------------------------------------------------------
 * qblocks - fast, easily-accessible, fully-decentralized data from blockchains
 * copyright (c) 2016, 2021 TrueBlocks, LLC (http://trueblocks.io)
 *
 * This program is free software: you may redistribute it and/or modify it under the terms
 * of the GNU General Public License as published by the Free Software Foundation, either
 * version 3 of the License, or (at your option) any later version. This program is
 * distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even
 * the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * General Public License for more details. You should have received a copy of the GNU General
 * Public License along with this program. If not, see http://www.gnu.org/licenses/.
 *-------------------------------------------------------------------------------------------*/
#include "etherlib.h"

//----------------------------------------------------------------
bool visitTransaction(CTransaction& trans, void* data) {
    trans.toJson(cout);
    return true;
}

//----------------------------------------------------------------
bool visitBlock(CBlock& block, void* data) {
    return block.forEveryTransaction(visitTransaction, data);
}

//----------------------------------------------------------------
int main(int argc, const char* argv[]) {
    loadEnvironmentPaths();

    etherlib_init(quickQuitHandler);

    int start = atoi(argv[1]);
    int nBlocks = atoi(argv[2]);

    cout << "Starting block: " << start << "\n";
    cout << "Num blocks " << nBlocks << "\n";

    uint32_t counter = 0;
    forEveryBlock(visitBlock, &counter, start, nBlocks, 1);

    etherlib_cleanup();

    return 0;
}
